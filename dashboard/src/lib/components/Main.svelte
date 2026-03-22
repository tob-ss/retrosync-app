<script lang="ts">
    import Sidebar from "./Sidebar.svelte";
    import Content from "./Content.svelte";
    import Login from "./Login.svelte";
    let loggedIn = $state(false);
    
</script>


<div class="h-full w-full">
    {#if loggedIn}
    <div class="h-full grid grid-cols-5 grid-rows-5 gap-4">
        <div class="row-span-5">
            <Sidebar />
        </div>
        <div class="col-span-4 row-span-5">
            <Content />
        </div>

    </div>
    {:else if loggedIn === false}
        <Login />
    {/if}
</div>

<style>


</style>
