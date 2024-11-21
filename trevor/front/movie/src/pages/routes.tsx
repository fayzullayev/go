import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom";
import Error from "./error";
import Home from "./home";
import Login from "./login";
import MainLayout from "../components/layout/main-layout";

export const routes = createBrowserRouter(
  createRoutesFromElements(
    <Route errorElement={<Error />}>
      <Route path={"/"} element={<MainLayout />}>
        <Route index={true} element={<Home />} />
        <Route path={"/login"} element={<Login />} />
      </Route>
    </Route>,
  ),
);
