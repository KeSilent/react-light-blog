"use client";

import { cn } from "@/lib/utils";

export default function Button({
  title,
  className = "",
  type = "button",
  onClick,
}: {
  title: string;
  onClick: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  className?: string;
  type?: "button" | "submit" | "reset";
}) {
  return (
    <button
      type={type}
      onClick={onClick}
      className={cn(
        "flex w-35 justify-center rounded-md bg-background px-3 py-1.5 text-sm/6 font-semibold text-foreground shadow-sm hover:bg-hover-background focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-background",
        className
      )}
    >
      {title}
    </button>
  );
}
