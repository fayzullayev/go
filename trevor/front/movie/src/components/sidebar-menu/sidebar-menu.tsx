import { SidebarMenuContainer } from "./style.ts";
import { AppstoreOutlined } from "@ant-design/icons";
import { Menu, MenuProps } from "antd";

type MenuItem = Required<MenuProps>["items"][number];

const items: MenuItem[] = [
  {
    key: "sub1",
    label: "Home",
    icon: <AppstoreOutlined />,
  },
  {
    key: "Movies",
    label: "Movies",
    icon: <AppstoreOutlined />,
  },
  {
    key: "Genres",
    label: "Genres",
    icon: <AppstoreOutlined />,
  },
  {
    key: "Add movie",
    label: "Add movie",
    icon: <AppstoreOutlined />,
  },
  {
    key: "Manage Catalogue",
    label: "Manage Catalogue",
    icon: <AppstoreOutlined />,
  },
  {
    key: "GraphQL",
    label: "GraphQL",
    icon: <AppstoreOutlined />,
  },
];

function SidebarMenu() {
  const onClick: MenuProps["onClick"] = (e) => {
    console.log("click ", e);
  };

  return (
    <SidebarMenuContainer>
      <Menu
        onClick={onClick}
        style={{ width: 256 }}
        defaultSelectedKeys={["1"]}
        defaultOpenKeys={["sub1"]}
        mode="inline"
        items={items}
      />
    </SidebarMenuContainer>
  );
}

export default SidebarMenu;
