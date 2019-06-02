# Redux 理解与运用

### 为什么需要Redux ?

**组件通信更加容易：**

![组件通信](https://css-tricks.com/wp-content/uploads/2016/03/redux-article-3-03.svg)

### Redux特性

1. Single Source of Truth

2. 可预测性：`state + action = new state` 

3. 纯函数产生新的 `store`。保证了可预测性

---

## 理解 Store，Action，Reducer

### Action

Action 是把数据从应用传到 Store 的有效载荷，也是Store数据的唯一来源。

```jsx
{
   type:ADD_TODO,
   text:"hello, Add This Now! "
}
```

### Reducer

Reducer 指定应用状态的变化如何响应 Action 并发送到 Store，记住 actions 只是描述了*有事情发生了*这一事实，并没有描述应用如何更新 state。

```jsx
import {
  ADD_TODO,
  TOGGLE_TODO,
  SET_VISIBILITY_FILTER,
  VisibilityFilters
} from './actions'

...

function todoApp(state = initialState, action) {
  switch (action.type) {
    case SET_VISIBILITY_FILTER:
      return Object.assign({}, state, {
        visibilityFilter: action.filter
      })
    case ADD_TODO:
      return Object.assign({}, state, {
        todos: [
          ...state.todos,
          {
            text: action.text,
            completed: false
          }
        ]
      })
    default:
      return state
  }
}
```

### Store

我们学会了使用  [action](https://www.redux.org.cn/docs/basics/Actions.html)  来描述“发生了什么”，和使用  [reducers](https://www.redux.org.cn/docs/basics/Reducers.html)  来根据 action 更新 state 的用法。

**Store**  就是把它们联系到一起的对象。

- 维持应用的 state；
- 提供 [`getState()`](https://www.redux.org.cn/docs/api/Store.html#getState) 方法获取 state；
- 提供 [`dispatch(action)`](https://www.redux.org.cn/docs/api/Store.html#dispatch) 方法更新 state；
- 通过 [`subscribe(listener)`](https://www.redux.org.cn/docs/api/Store.html#subscribe) 注册监听器;
- 通过 [`subscribe(listener)`](https://www.redux.org.cn/docs/api/Store.html#subscribe) 返回的函数注销监听器。

```jsx
const store = createStore(reducer)
```
