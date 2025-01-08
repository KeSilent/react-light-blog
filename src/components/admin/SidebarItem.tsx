"use client";

import NavbarDataApi from "@/api/navbar-data-api";
import { cn } from "@/lib/utils";
import { NavbarLinkModel } from "@/model/navbar-link-model";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

export default function SidebarItem() {
  const pathname = usePathname();
  const [links, setLinks] = useState<NavbarLinkModel[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let mounted = true;
    const fetchData = async () => {
      try {
        setLoading(true);
        const list = await NavbarDataApi();
        if (mounted) {
          setLinks(list);
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
    <>
      {links.map(({ id, title, url, icon }: NavbarLinkModel) => {
        const isActive = pathname === url;
        return (
          <Link
            key={id}
            href={url}
            className={cn(
              "flex items-center gap-3 rounded-lg px-3 py-2 transition-all hover:text-primary",
              { "bg-muted text-primary ": isActive },
              { "text-muted-foreground ": !isActive }
            )}
          >
            {icon}
            {title}
          </Link>
        );
      })}
    </>
  );
}