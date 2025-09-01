import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';

export class State {
  public id = 0; // 分类ID
  public code = ''; // 代码
  public name = ''; // 名称
  public exchange = ''; // 交易所

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}

export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// 表单验证规则
export const rules = {
  code: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入代码',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '分类ID',
    componentProps: {
      placeholder: '请输入分类ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '分类ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '代码',
    key: 'code',
    align: 'left',
    width: -1,
  },
  {
    title: '名称',
    key: 'name',
    align: 'left',
    width: -1,
  },
  {
    title: '交易所',
    key: 'exchange',
    align: 'left',
    width: -1,
  },
];