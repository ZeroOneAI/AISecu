import 'react-pro-sidebar/dist/css/styles.css';
import React, { useEffect, useState } from "react";
import {
    ProSidebar,
    Menu,
    MenuItem,
    SidebarHeader,
    SidebarFooter,
    SidebarContent,
} from "react-pro-sidebar";
import { FaRegHeart } from "react-icons/fa";
import { FiHome, FiLogOut, FiLogIn } from "react-icons/fi";
import { RiPencilLine } from "react-icons/ri";
import "./Header.css";
import { Link } from 'react-router-dom';
import { AiFillSecurityScan } from "react-icons/ai";


function Header({ headerRef }) {
    return <div id="header">
        {/* collapsed props to change menu size using menucollapse state */}
        <ProSidebar ref={headerRef}>
            <SidebarHeader>
                <div className="logotext" style={{ padding: "20px" }}>
                    <AiFillSecurityScan /> AI Secu
                </div>
            </SidebarHeader>
            <SidebarContent>
                <Menu iconShape="square">
                    <MenuItem icon={<FiHome />}>Home<Link to="/" /></MenuItem>
                    <MenuItem icon={<FaRegHeart />}>Repository<Link to="/repository" /></MenuItem>
                    <MenuItem icon={<RiPencilLine />}>Account<Link to="/account" /></MenuItem>
                </Menu>
            </SidebarContent>
        </ProSidebar>
    </div>
}

export default Header;
