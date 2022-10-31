import { useEffect, useState } from "react";
import { Button, FloatingLabel, Form, Modal } from "react-bootstrap";

function NewAccountSubmitValidate(event) {
    event.preventDefault();
    const registryType = event.target[0].value;
    const accountId = event.target[1].value;
    const accountPassword = event.target[2].value;
    const accountNickname = event.target[3].value;

    if (registryType === "Select Registry" || accountId === "" || accountPassword === "") {
        return false;
    }

    const submit = async () => {
        const resp = await fetch("/api/account", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                "account_nickname": accountNickname,
                "account_password": accountPassword,
                "account_username": accountId,
                "registry_type": registryType,
            })
        })
        await resp.json();
        window.location.reload();
    }
    submit();
    return false;
}

function NewAccount() {
    const [modalShow, setModalShow] = useState(false);
    const [typeList, setTypeList] = useState([]);

    useEffect(() => {
        setTypeList(["dockerhub"])
    }, []);

    return (
        <>
            <Button variant='dark' onClick={() => setModalShow(true)}>Add New Account</Button>
            <Modal show={modalShow} onHide={() => setModalShow(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title id="create-repository-modal">Add New Account</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form onSubmit={NewAccountSubmitValidate}>
                        <div>
                            <h4>Registry Type</h4>
                            <Form.Select aria-label="type-select">
                                <option>Select Registry Type</option>
                                {
                                    typeList.map((type, index) => {
                                        return (<option key={index} value={type}>{type}</option>);
                                    })
                                }
                            </Form.Select>
                        </div>
                        <hr style={{ width: 0 }}></hr>
                        <FloatingLabel label="Account Id"><Form.Control type="textarea" placeholder='account id'></Form.Control></FloatingLabel>
                        <p></p>
                        <FloatingLabel label="Account Password"><Form.Control type="password" placeholder='account password'></Form.Control></FloatingLabel>
                        <p></p>
                        <FloatingLabel label="Account Nickname"><Form.Control type="textarea" placeholder='account nickname'></Form.Control></FloatingLabel>
                        <hr style={{ width: 0 }}></hr>
                        <Button type="submit">Add</Button>
                    </Form>
                </Modal.Body>
            </Modal>

        </>
    )
}

export default NewAccount;

