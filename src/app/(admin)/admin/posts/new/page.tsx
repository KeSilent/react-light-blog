"use client";

import MDEditor from "@uiw/react-md-editor";
import { SetStateAction, useState } from "react";
import rehypeSanitize from "rehype-sanitize";
import {
  getCommands,
  getExtraCommands,
} from "@uiw/react-md-editor/commands-cn";

import "./mdEditorCss.css";
import Button from "@/components/UI/button";
import UploadButton from "@/components/UI/upload-button";
import CheckboxTree, { TreeNode } from "@/components/UI/checkbox-tree";
import fm from "front-matter";
import SelectMenu from "@/components/UI/select-menus";

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

  //选中的状态
  const [checkedState, setCheckedState] = useState<string>("1");

  const stateList = [
    {
      id: "1",
      name: "已发布",
    },
    {
      id: "2",
      name: "草稿",
    },
    {
      id: "3",
      name: "已删除",
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

  // 上传图片文件
  const handleUpload = async (file: File) => {
    try {
      const formData = new FormData();
      formData.append("file", file);
      // 使用 fetch 发送文件到服务器
      const response = await fetch("/api/upload", {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error(`上传失败，状态码: ${response.status}`);
      }

      const result = await response.json();
      console.log("上传成功:", result);
      console.log("文件上传成功");
    } catch (error) {
      console.error("上传失败:", error);
    }
  };

  // 上传 Markdown 文件
  const handleUploadMDFile = async (file: File) => {
    try {
      // 使用 FileReader 读取文件内容
      const reader = new FileReader();
      reader.readAsText(file);
      reader.onload = (event) => {
        if (event.target && event.target.result) {
          let fileContent: string;
          if (typeof event.target.result === "string") {
            fileContent = event.target.result;
          } else {
            // 如果是 ArrayBuffer，可以尝试将其转换为字符串
            const uint8Array = new Uint8Array(
              event.target.result as ArrayBuffer
            );
            fileContent = new TextDecoder().decode(uint8Array);
          }
          // 使用 front-matter 解析文件内容
          const content = fm(fileContent);
          console.log("front matter:", content.attributes);
        }
      };
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
      <div className="container flex">
        <div className="flex-1">
          <div>
            <input
              type="text"
              name="posttitle"
              id="posttitle"
              className=" w-full p-2 my-4 border border-gray-300 rounded-md text-lg"
              placeholder="添加标题"
            />
          </div>
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
          <div className="flex flex-col items-center justify-center p-4 shadow-md w-full">
            <div className="w-full flex flex-col items-center justify-center">
              <UploadButton
                className="w-36 py-4"
                title={"设置图片"}
                onUploadComplete={handleUploadComplete}
                onUploadStart={handleUploadStart}
                onUpload={handleUpload}
              ></UploadButton>
              <div className="w-full flex flex-row justify-between items-center p-4">
                <div className="flex-1 pt-1">状态：</div>
                <div className="flex-1">
                  <SelectMenu
                    selected={checkedState}
                    menuList={stateList}
                    setSelected={setCheckedState}
                  />
                </div>
              </div>
            </div>
            <div className="w-full h-px bg-gray-300 my-6"></div>
            <div className="w-full">
              <label>分类目录</label>
              <div className="border border-gray-300 rounded-md mt-4 p-2 overflow-y-auto max-h-64">
                <CheckboxTree treeData={nodes} />
              </div>
            </div>
            <div className="w-full h-px bg-gray-300 my-6"></div>
            <div className="w-full">
              <label>添加摘要</label>
              <textarea
                name="postsummary"
                id="postsummary"
                className="w-full p-2 my-4 border border-gray-300 rounded-md"
                placeholder="添加摘要"
              ></textarea>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
