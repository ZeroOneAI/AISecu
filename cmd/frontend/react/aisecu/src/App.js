import 'bootstrap/dist/css/bootstrap.min.css';
import Header from "./Header";
import { BrowserRouter, Route, Routes, Outlet } from "react-router-dom";
import { useEffect, useRef, useState } from 'react';
import { Col, Container, Row } from 'react-bootstrap';
import Dashboard from './Dashboard';
import RepositoryList from './repository/list';
import RepositoryDetail from './repository/detail';
import NewRepository from './repository/new';
import ImageDetail from './image/detail';
import AccountList from './account/list';
import AccountDetail from './account/detail';
import NewAccount from './account/new';


function DefaultPage() {
  const headerRef = useRef(null);
  const [width, setWidth] = useState(270);
  useEffect(() => {
    setWidth(headerRef.current.clientWidth);
  }, []);
  return (<div style={{ padding: "0px 0px 0px " + width + "px" }}>
    <Header headerRef={headerRef}></Header>
    <Outlet />
  </div>);
}


function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<DefaultPage />}>
            <Route path="repository">
              <Route path=":repositoryId" element={<RepositoryDetail />} />
              <Route index element={<Container>
                <Row><Col><h1>Repository List</h1></Col></Row>
                <Row><RepositoryList /></Row>
                <Row style={{ padding: "0 5vw 0 5vw" }}><NewRepository></NewRepository></Row>
              </Container>} />
            </Route>
            <Route path="account">
              <Route path=":accountId" element={<AccountDetail />} />
              <Route index element={<Container>
                <Row><Col><h1>Account List</h1></Col></Row>
                <Row><AccountList /></Row>
                <Row style={{ padding: "0 5vw 0 5vw" }}><NewAccount></NewAccount></Row>
              </Container>} />
            </Route>
            <Route path="image">
              <Route path=":imageId" element={<ImageDetail />} />
            </Route>
            <Route path="setting">
              <Route index element={<div>setting</div>} />
            </Route>
            <Route index element={<Dashboard />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
