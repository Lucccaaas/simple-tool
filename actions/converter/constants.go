package converter

var publicClassTemplate string = `@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class $className {
$body
}
`

var innerClassTemplate string = `
@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public static class $className {
$body
}`
