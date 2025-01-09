"use client";

import NavbarDataApi from "@/api/navbar-data-api";
import { cn } from "@/lib/utils";
import { NavbarLinkModel } from "@/model/navbar-link-model";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

function SidebarItem({
  items,
  pathname,
}: {
  items: NavbarLinkModel[];
  pathname: string;
}) {
  return (
    <>
      {items.map(({ id, title, url, icon }: NavbarLinkModel) => {
        const isActive = pathname === url;
        return (
          <div
            key={id}
            className={cn(
              "flex dark:hover:text-black hover:text-primary hover:rounded-lg hover:bg-gray-200",
              { "text-primary dark:text-white": isActive },
              { "text-muted-foreground ": !isActive }
            )}
          >
            <div
              className={cn("my-2 pl-2", {
                "text-primary dark:text-white border-l-2": isActive,
              })}
            ></div>
            <Link
              key={id}
              href={url}
              className={cn(
                "flex items-center gap-3 rounded-lg px-3 py-2 transition-all cursor-pointer"
              )}
            >
              {icon}
              {title}
            </Link>
          </div>
        );
      })}
    </>
  );
}

export default function Sidebar() {
  const pathname = usePathname();
  const [items, setItems] = useState<NavbarLinkModel[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let mounted = true;
    const fetchData = async () => {
      try {
        setLoading(true);
        const list = await NavbarDataApi();
        if (mounted) {
          setItems(list);
        }
      } catch (err) {
        if (mounted) {
          setError("Failed to load data");
          console.error(err);
        }
      } finally {
        if (mounted) {
          setLoading(false);
        }
      }
    };

    fetchData();

    return () => {
      mounted = false;
    };
  }, []);

  if (loading) {
    return <div>加载中...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }
  return (
    <div className="hidden bg-muted/40 md:block">
      <div className="flex h-full max-h-screen flex-col gap-2">
        <div className="flex h-14 items-center border-b border-gray-800 px-4 lg:h-[60px] lg:px-5">
          <Link href="/" className="flex items-center gap-2 font-semibold">
            <span className="text-lg text-primary">ReactLightBlog</span>
          </Link>
        </div>
        <div className="flex-1">
          <nav className="grid items-start px-2 text-lg font-medium lg:px-4">
            <SidebarItem items={items} pathname={pathname}></SidebarItem>
          </nav>
        </div>
      </div>
    </div>
  );
}
