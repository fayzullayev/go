import { MainLayoutContainer } from "./style.ts";
import { Outlet } from "react-router-dom";
import { Button } from "antd";
import SidebarMenu from "../../sidebar-menu";

function MainLayout() {
  return (
    <MainLayoutContainer>
      <div style={{ display: "flex", justifyContent: "flex-end" }}>
        <Button type="primary">Login</Button>
      </div>
      <h1 style={{ padding: 0, margin: 0 }}>Go Watch a Movie</h1>
      <hr />
      <div style={{ display: "grid", gridTemplateColumns: "300px 1fr" }}>
        <SidebarMenu />
        <div>
          <Outlet />
        </div>
      </div>
    </MainLayoutContainer>
  );
}

export default MainLayout;
