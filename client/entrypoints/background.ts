export default defineBackground(() => {
  // ツールバーのアイコンをクリックしたら右サイドパネルが開くようにする
  browser.sidePanel
    ?.setPanelBehavior({ openPanelOnActionClick: true })
    .catch((err) => console.error(err));
});
