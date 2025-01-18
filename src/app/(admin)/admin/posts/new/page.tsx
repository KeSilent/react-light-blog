"use client";

import MDEditor from "@uiw/react-md-editor";
import { useState } from "react";
import rehypeSanitize from "rehype-sanitize";
import {
  getCommands,
  getExtraCommands,
} from "@uiw/react-md-editor/commands-cn";

import "./mdEditorCss.css";
import Button from "@/components/UI/button";
import UploadButton from "@/components/UI/upload-button";
import CheckboxTree, { TreeNode } from "@/components/UI/CheckboxTree";

export default function PostNew() {
  const [value, setValue] = useState("");
  const nodes: TreeNode[] = [
    {
      id: "1",
      label: "Node 1",
      children: [
        {
          id: "1-1",
          label: "Node 1-1",
        },
        {
          id: "1-2",
          label: "Node 1-2",
          children: [
            {
              id: "1-2-1",
              label: "Node 1-2-1",
            },
          ],
        },
      ],
    },
    {
      id: "2",
      label: "Node 2",
    },
  ];
  // 定义一个新的 onChange 回调函数
  const handleEditorChange = (value?: string) => {
    if (value !== undefined) {
      setValue(value);
    }
  };

  const handlePublish = () => {
    // 发布文章的逻辑
  };

  const handleUploadStart = () => {
    console.log("上传开始");
  };

  const handleUploadComplete = () => {
    console.log("上传完成");
  };

  const handleUpload = async (file: File) => {
    // 实际的上传逻辑，例如使用 fetch 或 axios
    const formData = new FormData();
    formData.append("file", file);

    try {
      await new Promise((resolve) => setTimeout(resolve, 2000)); // 模拟 2 秒的上传时间

      console.log("文件上传成功");
    } catch (error) {
      console.error("上传失败:", error);
    }
  };

  return (
    <>
      <div className="flex justify-end">
        <Button
          className="w-40"
          title={"发布"}
          onClick={handlePublish}
        ></Button>
      </div>
      <div>
        <input
          type="text"
          name="posttitle"
          id="posttitle"
          className=" w-full p-2 border border-gray-300 rounded-md text-lg"
          placeholder="添加标题"
        />
      </div>
      <div className="container flex">
        <div className="flex-1">
          <MDEditor
            height={800}
            value={value}
            visibleDragbar={false}
            preview="edit"
            onChange={handleEditorChange}
            commands={[...getCommands()]}
            extraCommands={[...getExtraCommands()]}
            previewOptions={{
              rehypePlugins: [[rehypeSanitize]],
            }}
          />
        </div>
        <div className="flex-none pl-4">
          <div className="flex flex-col shadow-md w-full">
            <UploadButton
              className="w-36 py-4"
              title={"设置图片"}
              onUploadComplete={handleUploadComplete}
              onUploadStart={handleUploadStart}
              onUpload={handleUpload}
            ></UploadButton>

            <div className="pt-4">
              <label>分类目录</label>
              <div className="border border-gray-300 rounded-md mt-4 p-2 overflow-y-auto max-h-64">
                <CheckboxTree treeData={nodes} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
