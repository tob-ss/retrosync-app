<script lang="ts">
  import Sidebar from "./Sidebar.svelte";
  import Content from "./Content.svelte";
</script>


<div class="h-screen w-full">
    <div class="h-screen grid grid-cols-5 grid-rows-5 gap-4">
        <div class="row-span-5">
            <Sidebar />
        </div>
        <div class="col-span-4 row-span-5">
            <Content />
        </div>

    </div>
</div>

<style>


</style>
