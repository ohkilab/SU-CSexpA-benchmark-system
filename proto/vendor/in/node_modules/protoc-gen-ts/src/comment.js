"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.addDeprecatedJsDoc = void 0;
const ts = require("typescript");
function addDeprecatedJsDoc(node, deprecated) {
    if (deprecated) {
        ts.addSyntheticLeadingComment(node, ts.SyntaxKind.MultiLineCommentTrivia, "* @deprecated", true);
    }
    return node;
}
exports.addDeprecatedJsDoc = addDeprecatedJsDoc;
