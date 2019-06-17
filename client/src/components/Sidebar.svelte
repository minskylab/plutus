<script>
  import { fade, fly } from "svelte/transition";

  export let company = [
    { name: "Name", value: "Plutus Example" },
    { name: "Contact email", value: "contact@plutus.com" },
    { name: "Contact phone", value: "+51957821858" },
    { name: "Website", value: "plutus.io" },
    { name: "Currency", value: "PEN", options: ["PEN", "USD"] }
  ];

  export let bridge = {
    backend: "culqi",
    version: "0.0.3",
    status: "working"
  };

  export let edit = false;

  const onEditMode = () => {
    edit = edit ? false : true;
  };

  const saveChanges = () => {
    edit = edit ? false : true;
  };
</script>

<div class="bg-white shadow-md h-full">
  <div class="flex flex-col h-full p-4">
    <div class="flex justify-end">
      <div class="m-2 cursor-pointer">
        {#if !edit}
          <p
            class="text-md text-teal-500 select-none"
            on:click={onEditMode}
            in:fly={{ x: 10, duration: 300 }}
            out:fly={{ x: 10, duration: 300 }}>
            Edit
          </p>
        {/if}
      </div>

    </div>
    <div transition:fade class="w-full flex flex-col">
      <img
        class="mx-auto w-24 h-24 bg-red-200 rounded-full jutify-center flex mt-4"
        src="https://picsum.photos/200"
        alt="company logo" />
      <div class="mx-auto text-gray-600 py-6 px-2">Hello, XXXX</div>
    </div>

    <div class="w-full flex mt-4">
      <h1 class="text-gray-800 text-xl font-black mx-auto">Your Company</h1>
    </div>

    <div class="flex scroll flex-col my-4">
      {#each company as field}
        <div class="flex flex-col mt-4 px-4">
          <div class="text-gray-500 font-light text-sm"> {field.name}</div>
          {#if edit}
            {#if field.name === 'Currency'}
              <select
                class="block appearance-none bg-white focus:outline-0
                focus:shadow-outline border border-gray-300 rounded-lg py-1 px-4
                block w-full appearance-none leading-normal my-2"
                id="grid-state"
                bind:value={field.value}>
                {#each field.options as opt}
                  <option>{opt}</option>
                {/each}
              </select>
            {:else}
              <input
                class="bg-white focus:outline-0 focus:shadow-outline border
                border-gray-300 rounded-lg py-1 px-4 block w-full
                appearance-none leading-normal my-2"
                type="text"
                bind:value={field.value} />
            {/if}
          {:else}
            <div class="text-gray-800 my-2 text-md"> {field.value} </div>
          {/if}

        </div>
      {/each}
      {#if edit}
        <div
          in:fly={{ y: 10, duration: 300 }}
          out:fly={{ y: 10, duration: 300 }}
          class="flex bg-teal-400 py-2 rounded-lg cursor-pointer hover:shadow-md
          my-2 mx-4"
          on:click={saveChanges}>
          <div class="mx-auto select-none">Save changes</div>
        </div>
      {/if}
    </div>

    <div class="w-full flex mt-4">
      <h1 class="text-gray-800 text-xl font-black mx-auto">Your Bridge</h1>
    </div>

    <div class="w-full flex mt-4">
      <div class="mx-auto">
        <p class="text-center text-gray-700 font-lg">
           {`${bridge.backend} v${bridge.version}`}
        </p>
        <div class="flex items-center my-2">
          <div class="w-3 h-3 bg-green-400 rounded-full mx-1" />
          <div class="text-gray-500 capitalize">{bridge.status}</div>
        </div>
      </div>
    </div>
  </div>

</div>
