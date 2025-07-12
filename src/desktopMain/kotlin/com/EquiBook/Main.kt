package com.EquiBook

import androidx.compose.material.*
import androidx.compose.runtime.*
import androidx.compose.ui.*
import androidx.compose.ui.window.*
import androidx.compose.ui.input.key.*
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.*
import androidx.compose.foundation.Image
import androidx.compose.ui.unit.dp
import org.jetbrains.compose.resources.painterResource
import com.EquiBook.resources.Res
import com.EquiBook.resources.equibook_logo

@Composable
fun Logo() {
    Image(
        painter = painterResource(Res.drawable.equibook_logo),
        contentDescription = "EquiBook Logo",
        modifier = Modifier.size(64.dp)
    )
}

fun KeyBindsHandler(key: KeyEvent): Boolean {
    when (key.type) {
        // Add your key handling logic here
    }
    return true
}

fun main() {
    application {
        Window(
            onCloseRequest = ::exitApplication,
            title = "EquiBook",
            onKeyEvent = { key -> KeyBindsHandler(key) }
        ) {
            Column(
                modifier = Modifier.fillMaxSize(),
                horizontalAlignment = Alignment.CenterHorizontally,
                verticalArrangement = Arrangement.Center
            ) {
                Logo()
                Text("Welcome to EquiBook!")
            }
        }
    }
}
