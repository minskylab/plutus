<script>
  import { onMount, onDestroy } from "svelte";
  const onLogin = e => {
    if (credentials.username !== "" && credentials.username !== "") {
      alert(`${credentials.username}, ${credentials.password}`);
    }
  };

  let credentials = {
    username: "",
    password: ""
  };
</script>

<div class="w-full h-full">
  <div class="h-full flex bg-orange-100 p-0 m-0">
    <div
      class="flex mx-auto my-auto bg-white rounded-lg shadow-lg lg:w-1/4
      sm:w-1/2 md:-w-1/2 px-4 py-10">
      <div class="bg-white mx-auto">
        <div class="flex-1 text-gray-800 text-center px-4 py-2 m-2">
          Welcome to your Plutus Dashboard
        </div>
        <div class="flex-1 text-gray-500 text-left py-2 m-2">
          Your credentials
        </div>

        <div class="flex px-1 py-1 flex-col">
          <input
            class="bg-white focus:outline-0 focus:shadow-outline border
            border-gray-300 rounded-lg py-2 px-4 block w-full appearance-none
            leading-normal"
            type="text"
            bind:value={credentials.username}
            placeholder="username" />
          <input
            class="bg-white focus:outline-0 focus:shadow-outline border
            border-gray-300 rounded-lg py-2 px-4 block w-full appearance-none
            leading-normal my-4"
            type="password"
            bind:value={credentials.password}
            placeholder="password" />
          <a
            href="/dashboard"
            on:click={onLogin}
            class="bg-teal-500 hover:bg-teal-800 text-white font-bold py-2 px-6
            rounded-lg border-none mt-4 text-center w-1/2 ml-auto">
            ENTER
          </a>

        </div>
      </div>
    </div>
  </div>

</div>
