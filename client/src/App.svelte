<script>
  let image =
    "https://images.unsplash.com/photo-1586074299757-dc655f18518c?fit=crop&w=1268&q=80";

  function change() {
    console.log(image);
    chrome.storage.local.set({ background: image });
  }
  function windowToString() {
    return window.getSelection().toString();
  }

  chrome.tabs.query({ currentWindow: true, active: true }, function (tabs) {
    const tab = tabs[0];
    chrome.scripting.executeScript(
      {
        func: windowToString,
        target: {
          allFrames: true,
          tabId: tab.id,
        },
      },
      function (selection) {
        console.log(selection);
        document.getElementById("output").innerHTML = selection[0]?.result;
      }
    );
  });
</script>

<main>
  <div id="output" />
</main>
