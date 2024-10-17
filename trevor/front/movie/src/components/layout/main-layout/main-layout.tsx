import { MainLayoutContainer } from "./style.ts";
import { Outlet } from "react-router-dom";

function MainLayout() {
  return (
    <MainLayoutContainer>
      hr
      <Outlet />
    </MainLayoutContainer>
  );
}

export default MainLayout;
