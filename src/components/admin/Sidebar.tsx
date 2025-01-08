import Link from "next/link";
import SidebarItem from "./SidebarItem";

export default function Sidebar() {
  return (
    <div className="hidden bg-muted/40 md:block">
      <div className="flex h-full max-h-screen flex-col gap-2">
        <div className="flex h-14 items-center border-b border-gray-800 px-4 lg:h-[60px] lg:px-5">
          <Link href="/" className="flex items-center gap-2 font-semibold">
            <span className="text-lg text-primary">Echo English</span>
          </Link>
        </div>
        <div className="flex-1">
          <nav className="grid items-start px-2 text-lg font-medium lg:px-4">
            <SidebarItem></SidebarItem>
          </nav>
        </div>
      </div>
    </div>
  );
}
