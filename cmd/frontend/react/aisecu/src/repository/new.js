import { useEffect, useState } from "react";
import { Button, FloatingLabel, Form, Modal } from "react-bootstrap";

function NewRepositorySubmitValidate() {
    return (event) => {
        event.preventDefault();
        const newRepositoryName = event.target[0].value;
        const newRepositoryAccount = event.target[1].value;

        if (newRepositoryName === "") {
            console.log("invalid project name");
            return false;
        }
        if (newRepositoryAccount === "Select Account") {
            console.log("invalid repository account");
            return false;
        }

        const submit = async () => {
            const resp = await fetch("/api/repository", {
                method: "POST",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ repository_name: newRepositoryName, account_id: newRepositoryAccount }),
            })
            await resp.json();
            window.location.reload();
        }
        submit();
        return false;
    }
}

function NewRepository() {
    const [modalShow, setModalShow] = useState(false);
    const [accountList, setAccountList] = useState([]);

    useEffect(() => {
        const setAccount = async () => {
            const resp = await fetch("/api/account/list", { redirect: "follow" });
            switch (resp.status) {
                case 200:
                    const result = await resp.json();
                    setAccountList(result.accounts.map((val) => { return { id: val.id, type: val.registry_type, nickname: val.nickname } }));
                    break;
                default:
                    break;
            }
        }
        setAccount();
    }, []);

    return (
        <>
            <Button variant='dark' onClick={() => setModalShow(true)}>Add New Repository</Button>
            <Modal show={modalShow} onHide={() => setModalShow(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title id="create-repository-modal">Add New Repository</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form onSubmit={NewRepositorySubmitValidate()}>
                        <FloatingLabel label="Repository Name"><Form.Control type="textarea" placeholder='repository name'></Form.Control></FloatingLabel>
                        <hr style={{ width: 0 }}></hr>
                        <div>
                            <h4>Account</h4>
                            <Form.Select aria-label="account-select">
                                <option>Select Account</option>
                                {
                                    accountList.map((account, index) => {
                                        return (<option key={index} value={account.id}>{account.type}/{account.nickname}</option>);
                                    })
                                }
                            </Form.Select>
                        </div>
                        <hr style={{ width: 0 }}></hr>
                        <Button type="submit">Add</Button>
                    </Form>
                </Modal.Body>
            </Modal>

        </>
    )
}

export default NewRepository;
