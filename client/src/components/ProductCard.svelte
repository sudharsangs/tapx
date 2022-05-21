<script>
  export let productImageUrl;
  export let productUrl;
  export let productName;
  export let productPrice;
  export let productRating;
  const optional = (x) => x || "";
</script>

<a href={productUrl} target="_blank">
  <div class="product-card">
    <img src={productImageUrl} alt="product" class="product-image" />
    <div class="product-details">
      <h2>{optional(productName)}</h2>
      <p class="product-price">{`₹${productPrice}`}</p>
      <p class="product-rating">{optional(productRating)}</p>
    </div>
  </div>
</a>

<style>
  a,
  a:hover,
  a:focus,
  a:active {
    text-decoration: none;
    color: inherit;
  }
  .product-card {
    display: flex;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    transition: 0.3s;
    padding: 8px;
    border-radius: 4px;
    width: 100%;
    margin: 10px 0px;
  }
  .product-image {
    width: 100px;
  }
  .product-details {
    margin-left: 12px;
  }
  .product-price {
    font-size: 16px;
    font-weight: 600;
  }
  .product-rating {
    font-size: 14px;
    font-weight: 400;
  }
</style>
