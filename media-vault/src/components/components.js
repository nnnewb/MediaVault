import ChooseMedia from "@/components/ChooseMedia";

export default {
  install(app) {
    ChooseMedia._context = app._context;
  },
};
