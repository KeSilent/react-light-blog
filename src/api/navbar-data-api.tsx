import { NavbarLinkModel } from "@/model/navbar-link-model";
import { Home, MonitorPlay, Snail, Users } from "lucide-react";

//获取导航数据
export default function NavbarDataApi(): NavbarLinkModel[] {
  return [
    {
      id: "0",
      title: "首页",
      url: "/admin",
      icon: <Home className="h-5 w-5" />,
    },
    {
      id: "1",
      title: "文章管理",
      url: "/admin",
      icon: <Home className="h-5 w-5" />,
    },
    {
      id: "2",
      title: "网站管理",
      url: "/scene",
      icon: <Snail className="h-5 w-5" />,
      subItems: [
        {
          id: "2-1",
          title: "Profile",
          url: "/settings/profile",
        },
        {
          id: "2-2",
          title: "Account",
          url: "/settings/account",
        },
      ],
    },
    {
      id: "3",
      title: "用户管理",
      url: "/video",
      icon: <MonitorPlay className="h-5 w-5" />,
    },
    {
      id: "4",
      title: "权限管理",
      url: "/chat",
      icon: <Users className="h-5 w-5" />,
    },
  ];
}
