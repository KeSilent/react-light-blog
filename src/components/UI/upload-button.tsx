"use client";

import { cn } from "@/lib/utils";
import { useState } from "react";

export default function UploadButton({
  title,
  className = "",
  onUploadStart,
  onUploadComplete,
  onUpload,
}: {
  title: string;
  className?: string;
  onUploadStart?: () => void;
  onUploadComplete?: () => void;
  onUpload: (file: File) => Promise<void>;
}) {
  const [isUploading, setIsUploading] = useState(false);

  const handleFileChange = async (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    if (!file) return;

    if (onUploadStart) {
      onUploadStart();
    }
    setIsUploading(true);

    try {
      await onUpload(file);
      if (onUploadComplete) {
        onUploadComplete();
      }
    } catch (error) {
      console.error("上传失败:", error);
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <label
      htmlFor="dropzone-file"
      className={cn(
        "flex w-35 cursor-pointer justify-center rounded-md bg-background px-3 py-1.5 text-sm/6 font-semibold text-foreground shadow-sm hover:bg-hover-background focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-background",
        className
      )}
    >
      <div className="flex flex-col items-center justify-center">
        <p className="text-xs text-gray-500 dark:text-gray-400">{title}</p>
        {isUploading && <p className="text-xs text-blue-500">上传中...</p>}
      </div>
      <input
        id="dropzone-file"
        type="file"
        className="hidden"
        onChange={handleFileChange}
      />
    </label>
  );
}
