import { h, render } from "vue";
import ChooseMediaComponent from "./ChooseMedia.vue";

function choose_single() {
  return new Promise((resolve, reject) => {
    const container = document.createElement("div");
    const v_node = h(ChooseMediaComponent, {
      onConfirm: (choosed) => {
        container.remove();
        resolve(choosed);
      },
      onCancel: () => {
        container.remove();
        reject("cancelled");
      },
    });
    v_node.appContext = ChooseMediaComponent._context;
    render(v_node, container);
    document.body.appendChild(container);
  });
}

ChooseMediaComponent._context = null;
ChooseMediaComponent.choose_single = choose_single;

export default ChooseMediaComponent;
