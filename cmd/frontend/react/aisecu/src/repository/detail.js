import { useEffect, useState } from 'react';
import { Button, Modal} from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import Detail from '../Detail';
import { IoMdCopy } from "react-icons/io";

function getRepositoryDetail(setDetail, repositoryId) {
    const getDetail = async () => {
        const latestImageResp = await fetch("/api/repository/latest/" + repositoryId);
        switch (latestImageResp.status) {
            case 200:
                const latestImage = await latestImageResp.json();
                const resp = await fetch("/api/metrics/image/" + latestImage.image.id)
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
                            Type: "Repository",
                        });
                        break;
                    default:
                        break;
                }
                break;
            default:
                break;
        }
    }
    getDetail();
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
    Type: "Repository",
}

function GetMoreInfo(repositoryId) {

    const [modalShow, setModalShow] = useState(false);
    const [webhookUrl, setWebhookUrl] = useState("");

    useEffect(() => {
        setWebhookUrl("https://secu.cloud.ainode.ai/webhook/scan/" + repositoryId);
    }, [repositoryId]);

    return () => {

        return (
            <>
                <Button variant='dark' onClick={() => setModalShow(true)}>More</Button>
                <Modal show={modalShow} onHide={() => setModalShow(false)} centered>
                    <Modal.Header closeButton>
                        <Modal.Title id="create-project-modal">More Info</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <h3>Webhook Url</h3>
                        <p><span onClick={() => { navigator.clipboard.writeText(webhookUrl) }} ><IoMdCopy /></span>{webhookUrl}</p>
                    </Modal.Body>
                </Modal>
            </>
        )
    }
}

function deleteRespository(repositoryId) {
    return () => {
        const deleteFunc = async () => {
            const resp = await fetch("/api/repository/delete/" + repositoryId, {
                method: "DELETE",
                headers: { 'Content-Type': 'application/json' },
                body: "{}",
            });
            await resp.json();
            window.location.href = "/repository";
        }
        deleteFunc();
    }
}

export async function getImageTags(repositoryId) {
    const resp = await fetch("/api/repository/images/" + repositoryId);
    switch (resp.status) {
        case 200:
            const result = await resp.json();
            return result;
        default:
            break;
    }
    return { images: [{ id: "1", tag: "t" }, { id: "2", tag: "tt" }] }
}

async function getRepositoryInfo(repositoryId) {
    const resp = await fetch("/api/repository/detail/" + repositoryId);
    switch (resp.status) {
        case 200:
            const result = await resp.json();
            return result;
        default:
            break;
    }
}

function getRepositoryRelatedInfo(repositoryId, setRelatedInfo) {
    const _getRepositoryRelatedInfo = async () => {
        const images = await getImageTags(repositoryId);
        const repository = await getRepositoryInfo(repositoryId);
        setRelatedInfo({
            Type: "Tag",
            List: images.images.map((image) => { return { url: "/image/" + image.id, name: image.tag } }),
            Name: repository.repository.name,
        });
    }
    _getRepositoryRelatedInfo();
}

function RepositoryDetail() {
    const params = useParams();
    const [relatedInfo, setRelatedInfo] = useState({ Type: "", List: [] })
    const [detail, setDetail] = useState(defaultDetail);
    useEffect(() => {
        getRepositoryDetail(setDetail, params.repositoryId);
        getRepositoryRelatedInfo(params.repositoryId, setRelatedInfo);
    }, [params.repositoryId]);
    return (<Detail detail={detail} relatedInfo={relatedInfo} MoreInfo={GetMoreInfo(params.repositoryId)} deleteFunc={deleteRespository(params.repositoryId)} />)
}

export default RepositoryDetail;
