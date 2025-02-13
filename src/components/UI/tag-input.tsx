import React, { useState, useRef } from "react";

interface TagInputProps {
  initialTags?: string[];
  onTagsChange?: (tags: string[]) => void;
}

export default function TagInput({ initialTags = [], onTagsChange }: TagInputProps) {
  // 状态管理：标签列表和输入框的值
  const [tags, setTags] = useState<string[]>(initialTags);
  const [inputValue, setInputValue] = useState<string>("");

  // 使用 useRef 来获取输入框的引用
  const inputRef = useRef<HTMLInputElement | null>(null);

  // 处理输入框的按键事件
  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" || e.key === ",") {
      e.preventDefault(); // 阻止默认行为（如提交表单）
      const newTag = inputValue.trim();
      if (newTag && !tags.includes(newTag)) {
        setTags([...tags, newTag]); // 添加新标签
        setInputValue(""); // 清空输入框
        onTagsChange?.([...tags, newTag]);
      }
    }
  };

  // 删除标签
  const removeTag = (tagToRemove: string) => {
    const updatedTags = tags.filter((tag) => tag !== tagToRemove);
    setTags(updatedTags);
    onTagsChange?.(updatedTags);
  };

  // 处理输入框失去焦点事件
  const handleBlur = () => {
    const newTag = inputValue.trim();
    if (newTag && !tags.includes(newTag)) {
      setTags([...tags, newTag]); // 添加新标签
      setInputValue(""); // 清空输入框
      onTagsChange?.([...tags, newTag]);
    }
  };

  return (
    <div className="w-full">
      {/* 输入框 */}
      <input
        ref={inputRef}
        type="text"
        value={inputValue}
        onChange={(e) => setInputValue(e.target.value)}
        onKeyDown={handleKeyDown}
        onBlur={handleBlur}
        className="border border-gray-300 rounded-md  rounded p-2 w-full"
        placeholder="输入标签并回车"
      />

      {/* 标签容器 */}
      <div className="mt-2 flex flex-wrap gap-2">
        {tags.map((tag, index) => (
          <div
            key={tag} // 使用标签内容作为 key
            className="bg-blue-500 text-white rounded-full px-3 py-1 text-sm flex items-center"
          >
            {tag}
            <button
              onClick={() => removeTag(tag)}
              className="ml-2 text-xs font-bold"
            >
              ×
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}