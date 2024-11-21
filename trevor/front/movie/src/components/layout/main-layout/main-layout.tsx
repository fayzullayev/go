import { MainLayoutContainer } from "./style.ts";
import { Link, Outlet } from "react-router-dom";

function MainLayout() {
  return (
    <MainLayoutContainer>
      <div>
        <ul>
          <li>
            <Link to={"/"}>Home</Link>
          </li>
          <li>
            <Link to={"/movies"}>Movies</Link>
          </li>
          <li>
            <Link to={"/genres"}>Genres</Link>
          </li>
          <li>
            <Link to={"/add-movies"}>Add movie</Link>
          </li>
          <li>
            <Link to={"/manage-catalogue"}>Manage catalogue</Link>
          </li>
          <li>
            <Link to={"/graphql"}>Graphql</Link>
          </li>
        </ul>
      </div>
      <Outlet />
    </MainLayoutContainer>
  );
}

export default MainLayout;
