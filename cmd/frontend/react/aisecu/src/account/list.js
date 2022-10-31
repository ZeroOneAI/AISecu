import { Link } from 'react-router-dom';
import List from '../List';

function getElems(setFunc) {
    const _getElems = async () => {
        const resp = await fetch("/api/account/list", { redirect: "follow" });
        switch (resp.status) {
            case 200:
                const result = await resp.json();
                setFunc(result.accounts.map((val) => { return { id: val.id, type: val.registry_type, name: val.nickname } }));
                break;
            default:
                break;
        }
    }
    _getElems();
}

function AccountList() {
    return (
        <List
            getElems={getElems}
            elemToTableElem={(elem, index) => <tr key={elem.id}><td>{index}</td><td>{elem.type}</td><td><Link to={"/account/" + elem.id}>{elem.name}</Link></td></tr>}
            tableHeader={<tr><th>#</th><th>Type</th><th>Nickname</th></tr>} />);
}

export default AccountList;


