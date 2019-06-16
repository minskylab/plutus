<script>
  import { fade, fly } from "svelte/transition";

  export let data = {
    id: "90asdjas0d0a9sdj",
    state: "draft",
    created_at: "18:30 09/02/2019",
    updated_at: "18:36 09/02/2019",
    customer: { name: "Bregy Malpartida" },
    products: [
      {
        quantity: 1,
        name: "Dry Herb Vaporizer",
        detail: "lorem ipsuom lorem ",
        cost: { amount: 100, currency: "PEN", symbol: "s/" }
      }
    ]
  };
</script>

<div class="flex bg-white px-6 py-4 rounded-lg shadow-lg flex-col">
  <div class="flex flex-col w-full">
    <div class="flex justify-between">
      <div class="w-full flex flex-col">
        <h1 class="font-bold text-2xl">Sale</h1>
        <div class="text-gray-600 text-sm">{data.id}</div>
      </div>
      <div class="w-full flex flex-col">
        <h2 class="text-gray-600 text-sm text-right">State</h2>
        <div class="flex justify-end items-center">
          <div class="w-3 h-3 rounded-full bg-yellow-400 mr-2" />
          <div class="text-gray-800 text-lg text-right capitalize">
             {data.state}
          </div>
        </div>
      </div>
    </div>

    <div class="my-4">
      <div class="text-gray-900 text-lg font-bold">Products</div>
      {#each data.products as product}
        <div class="flex justify-between items-center mt-3">
          <div>
            <div class="text-md">{product.quantity} {product.name}</div>
            <div class="text-xs text-gray-500 max-w-xs truncate">
               {product.detail}
            </div>
          </div>

          <div class="flex items-baseline">
            <div class="text-gray-700 text-md mr-2 md:text-sm">
               {product.cost.symbol}
            </div>
            <div class="text-gray-800 text-xl md:text-md">
               {product.cost.amount.toFixed(2)}
            </div>
          </div>
        </div>
        <div class="my-1 h-1 w-full rounded-full bg-gray-100" />
      {/each}
    </div>
    <div class="flex justify-between items-center mt-3 font-bold">
      <div>
        <div class="text-md">TOTAL</div>
      </div>

      <div class="flex items-baseline">
        <div class="text-gray-700 text-md mr-2 md:text-sm"> {'s/'} </div>
        <div class="text-gray-800 text-xl md:text-md">
           {data.products
            .map(p => p.cost.amount)
            .reduce((total, i) => total + i)
            .toFixed(2)}
        </div>
      </div>
    </div>
  </div>
  <div class="flex w-full justify-between mt-6 mb-2 items-center">
    <div class="text-teal-400 text-left underline cursor-pointer">
      Print Invoice
    </div>
    <div class="px-6 py-2 bg-teal-400 rounded-lg cursor-pointer">Details</div>
  </div>
</div>
