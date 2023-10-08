*tempconv*
=====
Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan. 

Como usar?
----
Faça o download da biblioteca:

`go get https://github.com/ufla-gcc259/aula-git-parte-2@v1.0.1`

Pronto, agora é só usar:
```go
package main

import (
	"fmt"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	fmt.Printf("Que frio! %v\n", tempconv.AbsoluteZeroC) // Que frio! -273.15°C
	fmt.Printf("Fervendo! %v\n", tempconv.CToF(tempconv.BoilingC)) // Fervendo! 212°F
}
```

Outras linguagens?
----
Versões da biblioteca *tempconv* para outras linguagens:

> *Todo*


Licença
-----

> 
# Licença Apache, versão 2.0 Portugues.
Versão 2.0
Enviado: 8 de fevereiro de 2004
Remetente: Kevin Coar
Identificador curto SPDX: Apache-2.0
Comissário de bordo:Fundação de Software Apache
Link para a versão do administrador de licenças
Licença aprovada pela iniciativa de código aberto

## TERMOS E CONDIÇÕES DE USO, REPRODUÇÃO E DISTRIBUIÇÃO

### 1. Definições.
“Licença” significa os termos e condições de uso, reprodução e distribuição conforme definido nas Seções 1 a 9 deste documento.

“Licenciante” significa o proprietário dos direitos autorais ou entidade autorizada pelo proprietário dos direitos autorais que está concedendo a Licença.

“Pessoa Jurídica” significa a união da entidade atuante e todas as outras entidades que controlam, são controladas por ou estão sob controle comum com essa entidade. Para os fins desta definição, “controle” significa (i) o poder, direto ou indireto, de causar a direção ou gestão de tal entidade, seja por contrato ou de outra forma, ou (ii) propriedade de cinquenta por cento (50%) ou mais ações em circulação, ou (iii) propriedade beneficiária de tal entidade.

“Você” (ou “Seu”) significa uma pessoa física ou jurídica que exerce as permissões concedidas por esta Licença.

Formulário “Fonte” significa o formato preferido para fazer modificações, incluindo, entre outros, código-fonte de software, fonte de documentação e arquivos de configuração.

Formulário “Objeto” significa qualquer formato resultante de transformação mecânica ou tradução de um formulário Fonte, incluindo, entre outros, código-objeto compilado, documentação gerada e conversões para outros tipos de mídia.

“Trabalho” significa o trabalho de autoria, seja na forma de Fonte ou Objeto, disponibilizado sob a Licença, conforme indicado por um aviso de direitos autorais incluído ou anexado ao trabalho (um exemplo é fornecido no Apêndice abaixo).

“Trabalhos Derivados” significa qualquer trabalho, seja na forma Fonte ou Objeto, que seja baseado (ou derivado) do Trabalho e para o qual as revisões editoriais, anotações, elaborações ou outras modificações representem, como um todo, um trabalho original de autoria. Para os fins desta Licença, Trabalhos Derivados não incluirão trabalhos que permaneçam separáveis ​​ou meramente vinculados (ou vinculados pelo nome) às interfaces da Obra e dos Trabalhos Derivados da mesma.

“Contribuição” significa qualquer trabalho de autoria, incluindo a versão original do Trabalho e quaisquer modificações ou acréscimos a esse Trabalho ou Trabalhos Derivados do mesmo, que seja intencionalmente submetido ao Licenciador para inclusão no Trabalho pelo proprietário dos direitos autorais ou por um indivíduo ou Pessoa Jurídica autorizada a enviar em nome do proprietário dos direitos autorais. Para efeitos desta definição, “enviado” significa qualquer forma de comunicação electrónica, verbal ou escrita enviada ao Licenciante ou aos seus representantes, incluindo, entre outros, comunicação em listas de correio electrónico, sistemas de controlo de código fonte e sistemas de rastreio de problemas que são gerenciados por ou em nome do Licenciante com a finalidade de discutir e melhorar o Trabalho,

“Contribuidor” significa o Licenciante e qualquer pessoa física ou jurídica em nome da qual uma Contribuição tenha sido recebida pelo Licenciador e posteriormente incorporada na Obra.

### 2. Concessão de Licença de Direitos Autorais.
Sujeito aos termos e condições desta Licença, cada Colaborador concede a Você uma licença de direitos autorais perpétua, mundial, não exclusiva, gratuita, isenta de royalties e irrevogável para reproduzir, preparar Trabalhos Derivados de, exibir publicamente, executar publicamente, sublicenciar e distribuir a Obra e tais Obras Derivadas na forma de Fonte ou Objeto.

### 3. Concessão de Licença de Patente.
Sujeito aos termos e condições desta Licença, cada Colaborador concede a Você uma licença de patente perpétua, mundial, não exclusiva, gratuita, isenta de royalties e irrevogável (exceto conforme indicado nesta seção) para fazer, ter feito, usar, oferecer para vender, vender, importar e de outra forma transferir a Obra, onde tal licença se aplica apenas às reivindicações de patente licenciáveis ​​por tal Colaborador que são necessariamente infringidas apenas por sua(s) Contribuição(ões) ou pela combinação de sua(s) Contribuição(ões) com o Trabalho para o qual tal(is) Contribuição(ões) foi(ão) submetida(s). Se Você instaurar um litígio de patente contra qualquer entidade (incluindo uma reclamação cruzada ou reconvenção em uma ação judicial) alegando que a Obra ou uma Contribuição incorporada na Obra constitui violação de patente direta ou contributiva,

### 4. Redistribuição.
Você poderá reproduzir e distribuir cópias da Obra ou de Obras Derivadas dela em qualquer meio, com ou sem modificações, e na forma de Fonte ou Objeto, desde que atenda às seguintes condições:

Você deve fornecer a qualquer outro destinatário da Obra ou Obras Derivadas uma cópia desta Licença; e
Você deve fazer com que todos os arquivos modificados tenham avisos proeminentes informando que você alterou os arquivos; e
Você deve reter, na forma Fonte de quaisquer Obras Derivadas que Você distribuir, todos os avisos de direitos autorais, patentes, marcas registradas e atribuição da forma Fonte da Obra, excluindo aqueles avisos que não pertencem a qualquer parte das Obras Derivadas; e
Se a Obra incluir um arquivo de texto “AVISO” como parte de sua distribuição, então quaisquer Trabalhos Derivados que Você distribuir deverão incluir uma cópia legível dos avisos de atribuição contidos nesse arquivo AVISO, excluindo aqueles avisos que não pertencem a qualquer parte do Obras Derivadas, em pelo menos um dos seguintes locais: dentro de um arquivo de texto de AVISO distribuído como parte das Obras Derivadas; no formulário de origem ou na documentação, se fornecida juntamente com as Obras Derivadas; ou, dentro de uma exibição gerada pelas Obras Derivadas, se e onde tais avisos de terceiros normalmente aparecem. O conteúdo do arquivo AVISO é apenas para fins informativos e não modifica a Licença. Você pode adicionar Seus próprios avisos de atribuição nas Obras Derivadas que Você distribui, junto ou como um adendo ao texto do AVISO da Obra,
Você pode adicionar Sua própria declaração de direitos autorais às Suas modificações e pode fornecer termos e condições de licença adicionais ou diferentes para uso, reprodução ou distribuição de Suas modificações, ou para quaisquer Trabalhos Derivados como um todo, desde que Seu uso, reprodução e distribuição de o Trabalho, de outra forma, está em conformidade com as condições estabelecidas nesta Licença.

### 5. Envio de Contribuições.
A menos que Você declare explicitamente o contrário, qualquer Contribuição enviada intencionalmente para inclusão na Obra por Você ao Licenciante estará sob os termos e condições desta Licença, sem quaisquer termos ou condições adicionais. Não obstante o acima exposto, nada neste documento substituirá ou modificará os termos de qualquer contrato de licença separado que você possa ter celebrado com o Licenciador em relação a tais Contribuições.

### 6. Marcas registradas.
Esta Licença não concede permissão para usar nomes comerciais, marcas registradas, marcas de serviço ou nomes de produtos do Licenciador, exceto conforme necessário para uso razoável e habitual na descrição da origem da Obra e na reprodução do conteúdo do arquivo AVISO.

### 7. Isenção de Garantia.
A menos que exigido pela lei aplicável ou acordado por escrito, o Licenciante fornece o Trabalho (e cada Colaborador fornece suas Contribuições) “NO ESTADO EM QUE SE ENCONTRA”, SEM GARANTIAS OU CONDIÇÕES DE QUALQUER TIPO, expressas ou implícitas, incluindo, sem limitação, qualquer garantias ou condições de TÍTULO, NÃO VIOLAÇÃO, COMERCIALIZAÇÃO ou ADEQUAÇÃO A UM DETERMINADO FIM. Você é o único responsável por determinar a adequação do uso ou redistribuição da Obra e assume quaisquer riscos associados ao Seu exercício de permissões sob esta Licença.

### 8. Limitação de Responsabilidade.
Em nenhum caso e sob nenhuma teoria legal, seja por ato ilícito (incluindo negligência), contrato ou de outra forma, a menos que exigido pela lei aplicável (como atos deliberados e de negligência grave) ou acordado por escrito, qualquer Colaborador será responsável perante Você por danos, incluindo quaisquer danos diretos, indiretos, especiais, incidentais ou consequenciais de qualquer natureza decorrentes desta Licença ou do uso ou incapacidade de usar a Obra (incluindo, mas não se limitando a danos por perda de boa vontade, paralisação de trabalho , falha ou mau funcionamento do computador, ou todo e qualquer outro dano ou perda comercial), mesmo que tal Colaborador tenha sido avisado da possibilidade de tais danos.

### 9. Aceitação de Garantia ou Responsabilidade Adicional.
Ao redistribuir a Obra ou as Obras Derivadas da mesma, Você pode optar por oferecer e cobrar uma taxa pela aceitação de suporte, garantia, indenização ou outras obrigações de responsabilidade e/ou direitos consistentes com esta Licença. No entanto, ao aceitar tais obrigações, Você poderá agir apenas em Seu próprio nome e sob Sua exclusiva responsabilidade, não em nome de qualquer outro Colaborador, e somente se Você concordar em indenizar, defender e isentar cada Colaborador de qualquer responsabilidade incorrida por, ou reivindicações feitas contra tal Colaborador em razão de você aceitar qualquer garantia ou responsabilidade adicional.

## FIM DOS TERMOS E CONDIÇÕES

# Apache License, Version 2.0 Ingles
Version 2.0
Submitted: February 8, 2004
Submitter: Kevin Coar
SPDX short identifier: Apache-2.0
Steward:Apache Software Foundation
Link to license steward's version

## TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

### 1. Definitions.
“License” shall mean the terms and conditions for use, reproduction, and distribution as defined by Sections 1 through 9 of this document.

“Licensor” shall mean the copyright owner or entity authorized by the copyright owner that is granting the License.

“Legal Entity” shall mean the union of the acting entity and all other entities that control, are controlled by, or are under common control with that entity. For the purposes of this definition, “control” means (i) the power, direct or indirect, to cause the direction or management of such entity, whether by contract or otherwise, or (ii) ownership of fifty percent (50%) or more of the outstanding shares, or (iii) beneficial ownership of such entity.

“You” (or “Your”) shall mean an individual or Legal Entity exercising permissions granted by this License.

“Source” form shall mean the preferred form for making modifications, including but not limited to software source code, documentation source, and configuration files.

“Object” form shall mean any form resulting from mechanical transformation or translation of a Source form, including but not limited to compiled object code, generated documentation, and conversions to other media types.

“Work” shall mean the work of authorship, whether in Source or Object form, made available under the License, as indicated by a copyright notice that is included in or attached to the work (an example is provided in the Appendix below).

“Derivative Works” shall mean any work, whether in Source or Object form, that is based on (or derived from) the Work and for which the editorial revisions, annotations, elaborations, or other modifications represent, as a whole, an original work of authorship. For the purposes of this License, Derivative Works shall not include works that remain separable from, or merely link (or bind by name) to the interfaces of, the Work and Derivative Works thereof.

“Contribution” shall mean any work of authorship, including the original version of the Work and any modifications or additions to that Work or Derivative Works thereof, that is intentionally submitted to Licensor for inclusion in the Work by the copyright owner or by an individual or Legal Entity authorized to submit on behalf of the copyright owner. For the purposes of this definition, “submitted” means any form of electronic, verbal, or written communication sent to the Licensor or its representatives, including but not limited to communication on electronic mailing lists, source code control systems, and issue tracking systems that are managed by, or on behalf of, the Licensor for the purpose of discussing and improving the Work, but excluding communication that is conspicuously marked or otherwise designated in writing by the copyright owner as “Not a Contribution.”

“Contributor” shall mean Licensor and any individual or Legal Entity on behalf of whom a Contribution has been received by Licensor and subsequently incorporated within the Work.

### 2. Grant of Copyright License.
Subject to the terms and conditions of this License, each Contributor hereby grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare Derivative Works of, publicly display, publicly perform, sublicense, and distribute the Work and such Derivative Works in Source or Object form.

### 3. Grant of Patent License.
Subject to the terms and conditions of this License, each Contributor hereby grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable (except as stated in this section) patent license to make, have made, use, offer to sell, sell, import, and otherwise transfer the Work, where such license applies only to those patent claims licensable by such Contributor that are necessarily infringed by their Contribution(s) alone or by combination of their Contribution(s) with the Work to which such Contribution(s) was submitted. If You institute patent litigation against any entity (including a cross-claim or counterclaim in a lawsuit) alleging that the Work or a Contribution incorporated within the Work constitutes direct or contributory patent infringement, then any patent licenses granted to You under this License for that Work shall terminate as of the date such litigation is filed.

### 4. Redistribution.
You may reproduce and distribute copies of the Work or Derivative Works thereof in any medium, with or without modifications, and in Source or Object form, provided that You meet the following conditions:

You must give any other recipients of the Work or Derivative Works a copy of this License; and
You must cause any modified files to carry prominent notices stating that You changed the files; and
You must retain, in the Source form of any Derivative Works that You distribute, all copyright, patent, trademark, and attribution notices from the Source form of the Work, excluding those notices that do not pertain to any part of the Derivative Works; and
If the Work includes a “NOTICE” text file as part of its distribution, then any Derivative Works that You distribute must include a readable copy of the attribution notices contained within such NOTICE file, excluding those notices that do not pertain to any part of the Derivative Works, in at least one of the following places: within a NOTICE text file distributed as part of the Derivative Works; within the Source form or documentation, if provided along with the Derivative Works; or, within a display generated by the Derivative Works, if and wherever such third-party notices normally appear. The contents of the NOTICE file are for informational purposes only and do not modify the License. You may add Your own attribution notices within Derivative Works that You distribute, alongside or as an addendum to the NOTICE text from the Work, provided that such additional attribution notices cannot be construed as modifying the License.
You may add Your own copyright statement to Your modifications and may provide additional or different license terms and conditions for use, reproduction, or distribution of Your modifications, or for any such Derivative Works as a whole, provided Your use, reproduction, and distribution of the Work otherwise complies with the conditions stated in this License.

### 5. Submission of Contributions.
Unless You explicitly state otherwise, any Contribution intentionally submitted for inclusion in the Work by You to the Licensor shall be under the terms and conditions of this License, without any additional terms or conditions. Notwithstanding the above, nothing herein shall supersede or modify the terms of any separate license agreement you may have executed with Licensor regarding such Contributions.

### 6. Trademarks.
This License does not grant permission to use the trade names, trademarks, service marks, or product names of the Licensor, except as required for reasonable and customary use in describing the origin of the Work and reproducing the content of the NOTICE file.

### 7. Disclaimer of Warranty.
Unless required by applicable law or agreed to in writing, Licensor provides the Work (and each Contributor provides its Contributions) on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE. You are solely responsible for determining the appropriateness of using or redistributing the Work and assume any risks associated with Your exercise of permissions under this License.

### 8. Limitation of Liability.
In no event and under no legal theory, whether in tort (including negligence), contract, or otherwise, unless required by applicable law (such as deliberate and grossly negligent acts) or agreed to in writing, shall any Contributor be liable to You for damages, including any direct, indirect, special, incidental, or consequential damages of any character arising as a result of this License or out of the use or inability to use the Work (including but not limited to damages for loss of goodwill, work stoppage, computer failure or malfunction, or any and all other commercial damages or losses), even if such Contributor has been advised of the possibility of such damages.

### 9. Accepting Warranty or Additional Liability.
While redistributing the Work or Derivative Works thereof, You may choose to offer, and charge a fee for, acceptance of support, warranty, indemnity, or other liability obligations and/or rights consistent with this License. However, in accepting such obligations, You may act only on Your own behalf and on Your sole responsibility, not on behalf of any other Contributor, and only if You agree to indemnify, defend, and hold each Contributor harmless for any liability incurred by, or claims asserted against, such Contributor by reason of your accepting any such warranty or additional liability.

## END OF TERMS AND CONDITIONS


Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
