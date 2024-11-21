import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom";

import Error from "./error";
import Home from "./home";
import Login from "./login";
import MainLayout from "../components/layout/main-layout";

export const router = createBrowserRouter(
  createRoutesFromElements(
    <Route errorElement={<Error />} element={<MainLayout />}>
      <Route path={"/"}>
        <Route index={true} element={<Home />} />
        <Route path={"/movies"} element={<Home />} />
        <Route path={"login"} element={<Login />} />
      </Route>
    </Route>,
  ),
);
