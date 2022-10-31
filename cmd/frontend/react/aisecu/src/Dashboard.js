import { useEffect, useState } from "react";
import { Col, Container, Row, Table } from "react-bootstrap";
import { Link } from 'react-router-dom';

function Dashboard() {
    const [repositoryList, setRepositoryList] = useState([]);

    const update = async () => {
        const repositoryResp = await fetch("/api/repository", { redirect: "follow" });
        switch (repositoryResp.status) {
            case 200:
                const result = await repositoryResp.json();
                setRepositoryList(result.repositories.map((val) => { return { id: val.id, name: val.name } }));
                break;
            default:
                break;
        }
    };

    useEffect(() => {
        update()
    }, []);

    return <div>
        <Container style={{ textAlign: "center" }}>
            <Row style={{ height: "10vh" }}></Row>
            <Row>
                <Col>
                    <Table>
                        <thead>
                            <tr>
                                <th>Repository</th>
                            </tr>
                        </thead>
                        <tbody>
                            {
                                repositoryList.map((repository) => <tr key={repository.id}><td><Link to={"/repository/" + repository.id}>{repository.name}</Link></td></tr>)
                            }
                            <tr>
                                <td><Link to="/repository">{"<See More>"}</Link></td>
                            </tr>
                        </tbody>
                    </Table>
                </Col>
            </Row>
        </Container>
    </div>
}

export default Dashboard;
