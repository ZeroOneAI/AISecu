import { Link } from 'react-router-dom';
import List from '../List';

function getElems(setFunc) {
    const _getElems = async () => {
        const resp = await fetch("/api/repository", { redirect: "follow" });
        switch (resp.status) {
            case 200:
                const result = await resp.json();
                setFunc(result.repositories.map((val) => { return { id: val.id, name: val.name } }));
                break;
            default:
                break;
        }
    }
    _getElems();
}

function RepositoryList() {
    return (
        <List
            getElems={getElems}
            elemToTableElem={(elem, index) => <tr key={elem.id}><td>{index}</td><td><Link to={"/repository/" + elem.id}>{elem.name}</Link></td></tr>}
            tableHeader={<tr><th>#</th><th>Name</th></tr>} />);
}

export default RepositoryList;

