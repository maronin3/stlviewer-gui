import { LitElement, html } from "lit";
import { customElement } from "lit/decorators.js";

import "./component/MyHeader/myheader.ts";
import "./component/STLLoader/stlloder.ts";

@customElement("my-page")
class MyPage extends LitElement {
  render() {
    return html` 
      <my-header></my-header>
      <my-body></my-body>
    `;
  }
}
