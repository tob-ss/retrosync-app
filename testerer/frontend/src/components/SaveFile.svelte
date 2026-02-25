<script lang="ts">
    let game = $props();
    let altText = "Image of" + game.gameName
</script>

<div class="h-64 w-xl m-12 grid grid-cols-3 grid-rows-1 gap-4 p-12">
    <div class="col-span-2 col-start-2 row-start-1 grid grid-cols-3 grid-rows-3 gap-4 font-bold">
        <div class="col-span-3">
            <p>{game.gameName}</p>
        </div>
        <div class="col-span-3 row-start-2">
            <p>{game.console}</p>
        </div>
        <div class="row-start-3">
            <p>{game.device}</p>
        </div>
        <div class="col-start-3 row-start-3">
            <p>{game.timeMod}</p>
        </div>
        
        
        
    </div>
    <img class="col-start-1 row-start-1 h-full w-full object-contain object-center block" src={game.thumbnail} alt={altText} />
</div>