package storage

var PROMPTS = map[string]string{
	"NO_IMAGE": `Considerando o canal ou canais a serem analisados
	com o vídeo ou vídeos através do título ou títulos, descrição ou descrições, 
	e comentários abaixo, faça uma análise inteligente como um todo dos dados fornecidos
	com uma sugestão de vídeo que eu poderia fazer com 	base nos dados gerados a partir 
	do conteúdo fornecido. Abaixo o conteúdo a ser analisado:`,

	"IMAGE": `Considerando o canal ou canais a serem analisados
	com o vídeo ou vídeos através do título ou títulos, descrição ou descrições, 
	e comentários abaixo, faça uma análise inteligente como um todo dos dados fornecidos
	com uma sugestão de vídeo que eu poderia fazer com 	base nos dados gerados a partir 
	do conteúdo fornecido. Além disso, estou fornecendo a thumb ou thumbs dos vídeos; as imagens estão 
	na ordem respectivas ao número do vídeo fornecido - abaixo do nome do canal. Avalie as imagens fornecidas e os seus elementos
	para me dar sugestões de como eu poderia me basear para produção no tipo de conteúdo sugerido. Abaixo o conteúdo a ser analisado:`,
}
