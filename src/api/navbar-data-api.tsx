import { NavbarLinkModel } from "@/model/navbar-link-model";
import { Home, MonitorPlay, Snail, Users } from "lucide-react";

//获取导航数据
export default function NavbarDataApi(): NavbarLinkModel[] {
  return [
    {
      id: 1,
      title: "主页",
      url: "/",
      icon: <Home className="h-5 w-5" />,
    },
    {
      id: 2,
      title: "场景",
      url: "/scene",
      icon: <Snail className="h-5 w-5" />,
    },
    {
      id: 3,
      title: "视频",
      url: "/video",
      icon: <MonitorPlay className="h-5 w-5" />,
    },
    {
      id: 4,
      title: "对话",
      url: "/chat",
      icon: <Users className="h-5 w-5" />,
    },
  ];
}
