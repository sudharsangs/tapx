<script>
  import { writable } from "svelte/store";
  import ProductCard from "./components/ProductCard.svelte";
  const apiData = writable([]);
  const windowToString = () => {
    return window.getSelection().toString();
  };
  let searchTermContent;
  const getProductResults = (searchTerm) => {
    searchTermContent = searchTerm;
    const myHeaders = new Headers({
      "Access-Control-Allow-Origin": "no-cors",
    });
    fetch(`http://localhost:5600/api/v1/results?query=${searchTerm}`, {
      headers: myHeaders,
    })
      .then((response) => response.json())
      .then((data) => {
        apiData.set(data);
      })
      .catch((error) => {
        return [];
      });
  };

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
      (selection) => {
        getProductResults(selection[0]?.result);
      }
    );
  });
</script>

<main>
  <body style="width: 400px;height: 350px">
    <div class="product-list">
      {#each $apiData as data}
        <ProductCard
          productImageUrl={data.product_image_url}
          productName={data.product_name}
          productUrl={data.product_url}
          productPrice={data.product_price}
          productRating={data.product_rating}
        />
      {/each}
    </div>
  </body>
  <div class="button-wrapper">
    <a href={`https://www.amazon.in/s?k=${searchTermContent}`} target="_blank">
      <button>View More</button>
    </a>
  </div>
</main>

<style>
  a,
  a:hover,
  a:focus,
  a:active {
    text-decoration: none;
    color: inherit;
  }
  .product-list {
    display: flex;
    flex-direction: column;
  }
  .button-wrapper {
    position: fixed;
    bottom: 0px;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #cad5e2;
    padding: 8px;
  }
  button {
    background-color: #008cba;
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    border-radius: 4px;
    cursor: pointer;
  }
  button:hover {
    background-color: #008cba;
    border: none;
    color: white;
  }
</style>
