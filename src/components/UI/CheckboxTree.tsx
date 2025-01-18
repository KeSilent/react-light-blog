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
      <div key={node.id} style={{ marginLeft: 20 }}>
        <label>
          <input
            type="checkbox"
            checked={!!checkedState[node.id]}
            onChange={(e) => handleCheckboxChange(node, e.target.checked)}
          />
          {node.label}
        </label>
        {node.children && <div>{renderTreeNodes(node.children)}</div>}
      </div>
    ));
  };

  return <div>{renderTreeNodes(treeData)}</div>;
};

export default CheckboxTree;
