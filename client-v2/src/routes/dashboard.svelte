<script>
  import SaleCard from "../components/Sale.svelte";
  import Sidebar from "../components/Sidebar.svelte";

  let width = 0;
</script>

<svelte:window bind:innerWidth={width} />
<!-- 
{#if window !== undefined}
  <div class="absolute bg-red-500 right-0 bottom-0">
    w: {window.innerWidth} h: {window.innerHeight}
  </div>
{/if} -->

{#if width === 0}
  <div
    class="w-full h-full flex justify-center align-middle items-middle
    bg-orange-100">
    <div class="mx-auto my-auto">
      <h1>Loading...</h1>
    </div>
  </div>
{:else}
  <div class="h-full flex bg-orange-100 p-0 m-0">

    {#if width > 600}
      <div class="lg:w-2/12 h-full">
        <Sidebar />
      </div>
      <div class="lg:w-10/12 p-4">
        <!-- Helo -->
        <div class="flex w-full">
          <div class="w-1/3">
            <SaleCard />
          </div>
        </div>

      </div>
    {:else}
      <div class="lg:w-full h-full">
        <div class="flex w-full">

          <SaleCard />

        </div>
      </div>
    {/if}

  </div>

{/if}
