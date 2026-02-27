<script lang="ts">
 
</script>


<div class="h-screen w-full">
    <div class="h-screen bg-[#0f1118] grid grid-cols-1 grid-rows-8 gap-4 text-lg font-medium text-[#D7D6FC] font-heebo pr-12">
        <div>

        </div>
        <div class="row-span-3 col-start-1 row-start-2 ">
            <div class="flex flex-col justify-items-start align-items-center gap-12 pl-12">
                <h1>Saves</h1>
                <h1>Games</h1>
                <h1>Devices</h1>
            </div>
        </div>
            
        <div class="row-span-3 col-start-1 row-start-5">
            <div class="flex flex-col justify-items-start align-items-center gap-12 pl-12 pt-12">
                <h1>Memory Cards</h1>
                </div>
            <div class="flex grow flex-col place-items-center gap-12 pt-12 pl-12">
                <button class="rounded-xl border border-(--glass-border) w-full px-12 py-3 font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">+</button>
            </div>
        </div>
        <div class="col-start-1 row-start-8">

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
