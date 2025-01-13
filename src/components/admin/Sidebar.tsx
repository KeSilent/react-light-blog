"use client";

import NavbarDataApi from "@/api/navbar-data-api";
import { cn } from "@/lib/utils";
import { NavbarLinkModel } from "@/model/navbar-link-model";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { JSX, ReactNode, useEffect, useState } from "react";
//每个菜单项
function SidebarItem(
  id: string,
  isActive: boolean,
  url: string,
  icon: ReactNode,
  title: string
) {
  return (
    <div
      key={id}
      className={cn(
        "flex hover:text-primary hover:rounded-xl hover:bg-hover-background",
        { "text-primary": isActive }
      )}
    >
      <div
        className={cn("my-2 pl-2 border-l-4 border-primary", {
          "border-background": !isActive,
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
}

//每个菜单组
function SidebarGroup(
  id: string,
  isActive: boolean,
  icon: ReactNode,
  title: string,
  subItems: NavbarLinkModel[],
  expandedItems: Set<string>,
  onToggle: (id: string | null) => void
) {
  debugger
  const isExpanded = expandedItems.has(id);
  return (
    <div key={id}>
      <div
        className={cn(
          "flex gap-3 hover:text-primary hover:rounded-xl hover:bg-hover-background",
          { "text-primary": isActive }
        )}
      >
        <div
          className={cn("my-2 border-l-4 border-primary", {
            "border-background": !isActive,
          })}
        ></div>
        <div
          className="flex items-center gap-3 rounded-lg px-3 py-2 transition-all cursor-pointer"
          onClick={() => onToggle(isExpanded ? null : id)}
        >
          {icon}
          {title}
          <span className="ml-auto">
            {/* 添加展开/折叠图标 */}
            {isExpanded ? "▲" : "▼"}
          </span>
        </div>
      </div>
      <div className={cn("grid gap-2 ml-10", { hidden: !isExpanded })}>
        {subItems.map(({ id, title, url }: NavbarLinkModel) => {
          return SidebarItem(id, isActive, url, null, title);
        })}
      </div>
    </div>
  );
}
function SidebarItems({
  items,
  pathname,
  expandedItemId,
  onToggle,
}: {
  items: NavbarLinkModel[];
  pathname: string;
  expandedItemId: string | null;
  onToggle: (id: string | null) => void;
}) {
  const resultItems: JSX.Element[] = [];

  items.map(({ id, title, url, icon, type, subItems }: NavbarLinkModel) => {
    const isActive = pathname === url;
    if (type === "group") {
      resultItems.push(
        SidebarGroup(
          id,
          isActive,
          icon,
          title,
          subItems || [],
          new Set(expandedItemId ? [expandedItemId] : []), // 过滤掉 null
          onToggle
        )
      );
    } else if (type === "link") {
      resultItems.push(SidebarItem(id, isActive, url, icon, title));
    }
  });

  return <>{resultItems}</>;
}

export default function Sidebar() {
  const pathname = usePathname();
  const [items, setItems] = useState<NavbarLinkModel[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [expandedItemId, setExpandedItemId] = useState<string | null>(null); // 新增状态

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
        <div className="flex h-14 items-center border-b border-colos px-4 lg:h-[60px] lg:px-5">
          <Link href="/" className="flex items-center gap-2 font-semibold">
            <span className="text-lg text-primary">ReactLightBlog</span>
          </Link>
        </div>
        <div className="flex-1">
          <nav className="grid items-start px-2 text-xm font-medium lg:px-4">
            <SidebarItems
              items={items}
              pathname={pathname}
              expandedItemId={expandedItemId}
              onToggle={setExpandedItemId}
            ></SidebarItems>
          </nav>
        </div>
      </div>
    </div>
  );
}
