import { useEffect, useState } from 'react';
import { Col, Container, Row, Table } from 'react-bootstrap';

function List({ getElems, elemToTableElem, tableHeader }) {
    const [elems, setElems] = useState([]);

    useEffect(() => {
        getElems(setElems);
    }, [getElems]);
    return (
        <>
            <Container>
                <Row>
                    <Col>
                        <Table striped>
                            <thead>
                                {
                                    tableHeader
                                }
                            </thead>
                            <tbody>
                                {elems.map((elem, index) => {
                                    return elemToTableElem(elem, index);
                                })}
                            </tbody>
                        </Table>
                    </Col>
                </Row>
            </Container>
        </>
    );
}

export default List;
