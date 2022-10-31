import { useEffect, useState } from 'react';
import { Button, Col, Container, FloatingLabel, Form, Modal, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';

function updateNickname(accountId) {
    return (event) => {
        event.preventDefault();
        const nickname = event.target[0].value;
        if (nickname === "") {
            return false;
        }
        const _updateNickname = async () => {
            const resp = await fetch("/api/account/nickname/" + accountId, {
                method: "PUT",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ nickname: nickname }),
            })
            await resp.json();
            window.location.reload();
        }
        _updateNickname();
        return false;
    }
}

function UpdateNickname({ accountId }) {
    const [modalShow, setModalShow] = useState(false);

    return <>
        <Button variant='dark' onClick={() => setModalShow(true)}>Change Nickname</Button>
        <Modal show={modalShow} onHide={() => setModalShow(false)} centered>
            <Modal.Header closeButton>
                <Modal.Title id="update nickname">Update Nickname</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={updateNickname(accountId)}>
                    <FloatingLabel label="Nickname"><Form.Control type="textarea" placeholder='nickname'></Form.Control></FloatingLabel>
                    <p></p>
                    <Button type='submit'>Update</Button>
                </Form>
            </Modal.Body>
        </Modal>
    </>
}

function updatePassword(accountId) {
    return (event) => {
        event.preventDefault();
        const password = event.target[0].value;
        if (password === "") {
            return false;
        }
        const _updatePassword = async () => {
            const resp = await fetch("/api/account/password/" + accountId, {
                method: "PUT",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ password: password }),
            })
            await resp.json();
        }
        _updatePassword();
        return false;
    }
}

function UpdatePassword({ accountId }) {
    const [modalShow, setModalShow] = useState(false);

    return <>
        <Button variant='dark' onClick={() => setModalShow(true)}>Change Password</Button>
        <Modal show={modalShow} onHide={() => setModalShow(false)} centered>
            <Modal.Header closeButton>
                <Modal.Title id="update password">Update Password</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={updatePassword(accountId)}>
                    <FloatingLabel label="Password"><Form.Control type="password" placeholder='password'></Form.Control></FloatingLabel>
                    <p></p>
                    <Button type='submit'>Update</Button>
                </Form>
            </Modal.Body>
        </Modal>
    </>
}

function deleteFunc(accountId) {
    return () => {
        const deleteFunc = async () => {
            const resp = await fetch("/api/account/delete/" + accountId, {
                method: "DELETE",
                headers: { 'Content-Type': 'application/json' },
                body: "{}",
            });
            await resp.json();
            window.location.href = "/account";
        }
        deleteFunc();
    }
}

function AccountDetail() {
    const [accountInfo, setAccountInfo] = useState({});
    let params = useParams();
    const accountId = params.accountId;

    useEffect(() => {
        const _setInfo = async () => {
            const resp = await fetch("/api/account/detail/" + accountId);
            if (resp.status !== 200) {
                return;
            }
            const result = await resp.json();
            setAccountInfo({
                name: result.account.username,
                nickname: result.account.nickname,
                registryType: result.account.registry_type,
            })
        }
        _setInfo();
    }, [accountId]);

    return (<Container>
        <Row style={{ height: "5vh" }} />
        <Row>
            <Col>
                <h1>Account / {accountInfo.nickname}</h1>
            </Col>
        </Row>
        <Row style={{ height: "5vh" }} />
        <Row>
            <h5>Registry Type</h5><p>: {accountInfo.registryType}</p>
            <h5>Username</h5><p>: {accountInfo.name}</p>
            <h5>Nickname</h5><p>: {accountInfo.nickname}</p>
            <p><UpdateNickname accountId={accountId}></UpdateNickname></p>
            <p><UpdatePassword accountId={accountId}></UpdatePassword></p>
        </Row>
        <Row style={{ margin: "10px" }}>
            <Button variant='danger' onClick={deleteFunc(accountId)}>Delete</Button>
        </Row>
    </Container>);
}

export default AccountDetail;


