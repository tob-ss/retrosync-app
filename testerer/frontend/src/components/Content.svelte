<script lang="ts">
    import GamesList from "./GamesList.svelte";
</script>


<div class="h-full w-full">
    <div class="h-full grid grid-cols-1 grid-rows-6 gap-4">
        <div>Buttons</div>
        <div class="row-span-5">
            <div class="h-full justify-items-center grid grid-cols-1 grid-rows-7 gap-4">
                <div class="col-span-1 row-span-7 row-start-1">
                    <GamesList />
                </div>
            </div>
        </div>
    </div>
    
</div>

<style>


</style>
