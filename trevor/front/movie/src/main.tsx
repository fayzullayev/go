import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router-dom";
import { routes } from "./pages/routes.tsx";
import "./index.css";

createRoot(document.getElementById("root")!).render(
  <RouterProvider router={routes} />,
);
