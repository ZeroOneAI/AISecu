import { useEffect, useState } from 'react';
import { Button, Modal, Table } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import Detail from '../Detail';

function getImageDetail(setDetail, imageId) {
    const _setDetail = async () => {

        const resp = await fetch("/api/metrics/image/" + imageId)
        switch (resp.status) {
            case 200:
                const result = await resp.json();
                setDetail({
                    Info: {
                        Vulnerabilites: {
                            Critical: result.critical,
                            High: result.high,
                            Medium: result.medium,
                            Low: result.low,
                            Unassigned: result.unassigned,
                        },
                        PolicyViolations: {
                            LicenseRisk: result.policyViolationsLicenseTotal,
                            OperationalRisk: result.policyViolationsOperationalTotal,
                        }
                    },
                    Type: "Image",
                });
                break;
            default:
                break;
        }
    }
    _setDetail();
}

function getImageRelatedInfo(setRelatedInfo, imageId) {
    const _get = async () => {
        const imageResp = await fetch("/api/image/detail/" + imageId);
        if (imageResp.status !== 200) {
            return;
        }
        const imageResult = await imageResp.json();
        const repositoryResp = await fetch("/api/repository/detail/" + imageResult.image.repository_id);
        if (repositoryResp.status !== 200) {
            return;
        }
        const repositoryResult = await repositoryResp.json();
        setRelatedInfo({
            Name: repositoryResult.repository.name + ":" + imageResult.image.tag,
        });
    }
    _get();
}

const defaultDetail = {
    Info: {
        Vulnerabilites: {
            Critical: 0,
            High: 0,
            Medium: 0,
            Low: 0,
            Unassigned: 0,
        },
        PolicyViolations: {
            LicenseRisk: 0,
            OperationalRisk: 0,
        }
    },
    Type: "Image",
}

function CVEElem({ index, packageName, pacakgeVersion, cveID, cweName, cweDescription, cweId, severity }) {
    const [modalShow, setModalShow] = useState(false);

    return <tr>
        <td>{index}</td>
        <td>{packageName}</td>
        <td>{pacakgeVersion}</td>
        <td>
            <span style={{ color: "blue", textDecoration: "underline" }} onClick={() => setModalShow(true)}>{cveID}</span>
            <Modal size='lg' show={modalShow} onHide={() => setModalShow(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title id="cve-elem-detail">CVE Detail</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <h3>Vulnerability ID</h3>
                    <p>{cveID}</p>
                    <h3>Vulnerability Name</h3>
                    <p>{cweName}</p>
                    <h3>Vulnerability Description</h3>
                    <p>{cweDescription}</p>
                    {
                        cweId != null ? <a href={"https://cwe.mitre.org/data/definitions/" + cweId}>More</a> : null
                    }
                </Modal.Body>
            </Modal>
        </td>
        <td>{severity}</td>
    </tr>
}

function More(imageId) {
    const [modalShow, setModalShow] = useState(false);
    const [cveList, setCVEList] = useState([]);

    useState(() => {
        const getCVEList = async () => {
            const resp = await fetch("/api/cve/image/" + imageId);
            switch (resp.status) {
                case 200:
                    const result = await resp.json();
                    setCVEList(result);
                    break;
                default:
                    break;
            }
        }
        getCVEList();
    }, []);

    return () => {


        return <div style={{ margin: "10px" }}>
            <Button onClick={() => setModalShow(true)}>Look for Vulnerabilities Detail</Button>
            <Modal size='lg' show={modalShow} onHide={() => setModalShow(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title id="cve-detail">Vulnerabilities Detail</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <h3>{"<Vulnerabilities List>"}</h3>
                    <Table striped bordered hover>
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>Package Name</th>
                                <th>Pacakge Version</th>
                                <th>CVE</th>
                                <th>Severity</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                cveList.map((cve, index) => {
                                    return <CVEElem index={index} id={index} packageName={cve.component.name} pacakgeVersion={cve.component.version} cveID={cve.vulnerability.vulnId} cweId={cve.vulnerability.cweId} cweName={cve.vulnerability.cweName} cweDescription={cve.vulnerability.description} severity={cve.vulnerability.severity} />
                                })
                            }
                        </tbody>
                    </Table>
                </Modal.Body>
            </Modal>
        </div>
    }
}

function ImageDetail() {
    const params = useParams();
    const [relatedInfo, setRelatedInfo] = useState({ Type: "", List: [] });
    const [detail, setDetail] = useState(defaultDetail);
    const imageId = params.imageId;
    useEffect(() => {
        getImageDetail(setDetail, imageId);
        getImageRelatedInfo(setRelatedInfo, imageId);
    }, [imageId]);
    return (<Detail detail={detail} relatedInfo={relatedInfo} More={More(imageId)} />)
}

export default ImageDetail;
