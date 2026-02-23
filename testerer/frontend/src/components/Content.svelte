<script lang="ts">
    import GamesList from "./GamesList.svelte";
</script>


<div class="h-full w-full">
    <div class="grid grid-cols-1 grid-rows-6 gap-4">
        <div>Buttons</div>
        <div class="row-span-5">
            <div class="h-full grid grid-cols-3 grid-rows-7 gap-4">
                <div class="col-span-3">Header</div>
                <div class="col-span-3 row-span-6 row-start-2">
                    <GamesList />
                </div>
            </div>
        </div>
    </div>
    
</div>

<style>


</style>
