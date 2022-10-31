import React from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';
import { Button, Col, Container, Row, Table } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { FaBalanceScale } from 'react-icons/fa';
import { IoLayers } from 'react-icons/io5';

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
);

export const options = {
    indexAxis: 'y',
    elements: {
        bar: {
            borderWidth: 2,
        },
    },
    responsive: true,
    plugins: {
        title: {
            display: false,
            text: 'Vulnerabilities',
        },
    },
    scales: {
        x: {
            stacked: true,
        },
        y: {
            beginAtZero: true,
            stacked: true,
        },
    },
};

const labels = ['Vuln'];

const genData = (info) => {
    return {
        labels,
        datasets: [
            {
                label: 'Critical',
                data: [info.Vulnerabilites.Critical],
                borderColor: 'rgb(255, 0, 0)',
                backgroundColor: 'rgba(255, 0, 0, 0.5)',
            },
            {
                label: 'High',
                data: [info.Vulnerabilites.High],
                borderColor: 'rgb(255, 106, 0)',
                backgroundColor: 'rgba(255, 106, 0, 0.5)',
            },
            {
                label: 'Medium',
                data: [info.Vulnerabilites.Medium],
                borderColor: 'rgb(255, 242, 0)',
                backgroundColor: 'rgba(255, 242, 0, 0.5)',
            },
            {
                label: 'Low',
                data: [info.Vulnerabilites.Low],
                borderColor: 'rgb(34, 201, 0)',
                backgroundColor: 'rgba(34, 201, 0, 0.5)',
            },
            {
                label: 'Unassigned',
                data: [info.Vulnerabilites.Unassigned],
                borderColor: 'rgb(150, 150, 150)',
                backgroundColor: 'rgba(150, 150, 150, 0.5)',
            },
        ],
    }
};

function Detail({ detail, relatedInfo, MoreInfo, deleteFunc, AddRelated, DeleteRelated, More }) {
    return <Container>
        <Row style={{ height: "5vh" }} />
        <Row>
            <Col>
                <h1>{detail.Type} / {relatedInfo.Name}</h1>
            </Col>
            {MoreInfo != null ? <Col style={{ textAlign: "right" }}><MoreInfo /></Col> : null}
        </Row>
        <Row style={{ height: "5vh" }} />
        <Row>
            <Col>
                <Row>
                    <h3>{detail.Type} Vulnerabilites</h3>
                    <div style={{ width: "550px" }}><Bar options={options} data={genData(detail.Info)} /></div>
                </Row>
                <Row>
                    <Table>
                        <thead>
                            <tr>
                                <th colSpan={3}><h3>Policy Violations</h3></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td><FaBalanceScale /></td>
                                <td>License Risk</td>
                                <td>{detail.Info.PolicyViolations.LicenseRisk}</td>
                            </tr>
                            <tr>
                                <td><IoLayers /></td>
                                <td>Operational Risk</td>
                                <td>{detail.Info.PolicyViolations.OperationalRisk}</td>
                            </tr>
                        </tbody>
                    </Table>
                </Row>
            </Col>
            {
                relatedInfo != null && relatedInfo.Type != null && relatedInfo.Type !== "" ?
                    <Col>
                        <div>
                            <Table>
                                <thead>
                                    <tr>
                                        <th><h3>{relatedInfo.Type}</h3></th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        relatedInfo.List.map((elem, index) => {
                                            return (<tr key={index}><td><Link to={elem.url}>{elem.name}</Link> {DeleteRelated != null ? <DeleteRelated id={elem.id} /> : null}</td></tr>);
                                        })
                                    }
                                    {AddRelated != null ? <tr><td><AddRelated /></td></tr> : null}
                                </tbody>
                            </Table>
                        </div>
                    </Col> : null
            }
        </Row>
        {
            More != null ?
                <Row>
                    <More />
                </Row> : null
        }
        {
            deleteFunc != null ?
                <Row style={{ margin: "10px" }}>
                    <Button variant='danger' onClick={deleteFunc}>Delete</Button>
                </Row> : null
        }
    </Container>
}

export default Detail;
