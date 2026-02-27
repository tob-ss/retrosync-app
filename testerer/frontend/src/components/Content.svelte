<script lang="ts">
    import GamesList from "./GamesList.svelte";
</script>


<div class="h-full w-full">
    <div class="h-full grid grid-cols-2 grid-rows-6 gap-4">
        <div class="col-start-1 row-start-1 text-xl font-medium text-[#D7D6FC] font-heebo flex flex-col place-items-center p-12">
                <button class=" rounded-xl border border-(--glass-border) w-full max-w-md px-16 py-4 font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">Backup</button>
            </div>
        <div class="col-start-2 row-start-1 text-xl font-medium text-[#D7D6FC] font-heebo flex grow flex-col place-items-center p-12">
                <button class="rounded-xl border border-(--glass-border) w-full max-w-md px-16 py-4 font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">Download</button>
            </div>
        <div class="row-span-5 col-span-2 row-start-2">
            <div class="h-full justify-items-center grid grid-cols-1 grid-rows-7 gap-4">
                <div class="col-span-1 row-span-7 row-start-1">
                    <GamesList />
                </div>
            </div>
        </div>
    </div>
    
</div>

<style>
@import "tailwindcss";

button {
  --glass-bg: color-mix(in oklab, var(--bg) 80%, transparent);
  --glass-border: color-mix(in oklab, var(--glass-bg) 80%, rgb(255, 253, 120));
}

</style>
