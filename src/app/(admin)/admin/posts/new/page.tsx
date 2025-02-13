"use client";

import { useEffect, useState } from "react";
import Button from "@/components/UI/button";
import UploadButton from "@/components/UI/upload-button";
import CheckboxTree, { TreeNode } from "@/components/UI/checkbox-tree";
import SelectMenu from "@/components/UI/select-menus";
import CategoryDataApi from "@/api/category-data.api";
import { PostModel } from "@/model/post-model";
import TagInput from "@/components/UI/tag-input";
import { MdEditor } from 'md-editor-rt';
import 'md-editor-rt/lib/style.css';
import '../new/mdEditorCss.css'

export default function PostNew() {
  const [codeTheme] = useState('atom');
  const [categoryData, setCategoryData] = useState<TreeNode[]>([]);
  const [postModel, setPostModel] = useState<PostModel>({
    id: "", // 初始化必填字段
    title: "",
    slug: "",
    content: "",
    state: "2",
    cover: "",
    description: "",
    author: {
      id: "",
      name: "",
      email: "",
      avatar: ""
    },
    category: {
      id: "",
      name: "",
    },
    keys: []
  });

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

  useEffect(() => {
    const fetchData = async () => {
      setCategoryData(await CategoryDataApi());
    }
    fetchData();
  }, [])


  const handlePublish = () => {
    // 发布文章的逻辑
    console.log("发布文章", postModel);

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

  const onUploadImg = async (files: File[], callback: (arg0: any[]) => void) => {
    console.log("onUploadImg");

    // const res = await Promise.all(
    //   files.map((file) => {
    //     return new Promise((rev, rej) => {
    //       const form = new FormData();
    //       form.append('file', file);
    //     });
    //   })
    // );

    // callback(res.map((item) => item.data.url));
  };

  const findNodeById = (treeData: TreeNode[], id: string): TreeNode | undefined => {

    for (const node of treeData) {
      if (node.id === id) {
        return node;
      }
      if (node.children) {
        const foundNode = findNodeById(node.children, id);
        if (foundNode) {
          return foundNode;
        }
      }
    }
    return undefined;
  };


  // 编辑器回调函数
  const handleEditorChange = (value?: string) => {
    if (value !== undefined) {
      setPostModel(prev => ({
        ...prev,
        content: value,
      }));
    }
  };

  //设置状态回调函数
  const handleSelectedHandleChange = (value: { id: string; name: string }) => {
    setPostModel(prev => ({
      ...prev,
      state: value.id,
    }));
  };
  const [checkedState, setCheckedState] = useState<{ [key: string]: boolean }>({});
  useEffect(() => {
    const lastCheckedId = Object.keys(checkedState).reverse().find(id => checkedState[id]);
    if (lastCheckedId) {
      const selectCategory = findNodeById(categoryData, lastCheckedId);
      if (selectCategory) {
        setPostModel(prev => ({
          ...prev,
          category: {
            id: selectCategory.id,
            name: selectCategory.label,
          },
        }));
      }
    }
  }, [checkedState]);

  const handleCheckedChange = (newCheckedState: { [key: string]: boolean }) => {
    setCheckedState(newCheckedState);
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
              value={postModel?.title}
              onChange={(e) => setPostModel(prev => ({ ...prev, title: e.target.value }))}
            />
          </div>
          <MdEditor value={postModel?.content}
            onChange={handleEditorChange}
            onUploadImg={onUploadImg}
            preview={false}
            toolbarsExclude={["github"]}
            codeTheme={codeTheme}
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
                    selected={postModel?.state || ""}
                    menuList={stateList}
                    handleChange={handleSelectedHandleChange}
                  />
                </div>
              </div>
            </div>
            <div className="w-full h-px bg-gray-300 my-6"></div>
            <div className="w-full">
              <label>分类目录</label>
              <div className="border border-gray-300 rounded-md mt-4 p-2 overflow-y-auto max-h-64">
                <CheckboxTree treeData={categoryData} checkedState={checkedState} onCheckedChange={handleCheckedChange} />
              </div>
            </div>
            <div className="w-full h-px bg-gray-300 my-6"></div>
            <div className="w-full">
              <label>添加标签</label>
              <div className="mt-4 p-2 overflow-y-auto max-h-64">
                <TagInput initialTags={postModel?.keys} />
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
