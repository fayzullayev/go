import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
} from "react-router-dom";
import Error from "./error";
import Home from "./home";
import Login from "./login";

export const routes = createBrowserRouter(
  createRoutesFromElements(
    <Route errorElement={<Error />}>
      <Route path={"/"}>
        <Route index={true} element={<Home />} />
      </Route>
      <Route path={"/login"} element={<Login />} />
    </Route>,
  ),
);
