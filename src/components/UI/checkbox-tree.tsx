import React, { useState } from "react";

export interface TreeNode {
  id: string;
  label: string;
  children?: TreeNode[];
}

type CheckboxTreeProps = {
  treeData: TreeNode[];
};

const CheckboxTree: React.FC<CheckboxTreeProps> = ({ treeData }) => {
  const [checkedState, setCheckedState] = useState<{ [key: string]: boolean }>(
    {}
  );

  const handleCheckboxChange = (node: TreeNode, checked: boolean) => {
    setCheckedState((prevState) => ({
      ...prevState,
      [node.id]: checked,
    }));

    // 如果节点有子节点，递归更新子节点的状态
    if (node.children) {
      node.children.forEach((child) => {
        handleCheckboxChange(child, checked);
      });
    }
  };

  const renderTreeNodes = (nodes: TreeNode[]) => {
    return nodes.map((node) => (
      <div key={node.id} className="ml-3">
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            checked={!!checkedState[node.id]}
            onChange={(e) => handleCheckboxChange(node, e.target.checked)}
            className="form-checkbox h-4 w-4 text-blue-600"
          />
          <span className="text-gray-700">{node.label}</span>
        </label>
        {node.children && (
          <div className="ml-3 mt-2">{renderTreeNodes(node.children)}</div>
        )}
      </div>
    ));
  };

  return <div>{renderTreeNodes(treeData)}</div>;
};

export default CheckboxTree;
