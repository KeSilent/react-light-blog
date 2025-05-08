/*
 * @Author: Yang
 * @Date: 2025-04-22 09:39:37
 * @Description: 增加字段
 */
import { CodeBuilderFieldModel } from '@/models/system/code-builder-fields-model';
import { ActionType } from '@ant-design/pro-components';

export type CreateFieldProps = {
  model?: CodeBuilderFieldModel;
  reload?: ActionType['reload'];
};
export default function CreateField(props: CreateFieldProps) {
    return <div>CreateField</div>;
}
