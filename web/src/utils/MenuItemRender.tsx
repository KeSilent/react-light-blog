import * as AntdIcons from '@ant-design/icons';
import React from 'react';

export const createIcon = (icon: string | React.ReactNode) => {
  if (typeof icon === 'string') {
    const IconComponent = (AntdIcons as any)[icon];
    return IconComponent ? React.createElement(IconComponent) : null;
  }
  return icon;
};