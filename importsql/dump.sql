-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Feb 24, 2026 at 08:27 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `plan2go`
--

-- --------------------------------------------------------

--
-- Table structure for table `email_verification`
--

CREATE TABLE `email_verification` (
  `id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `otp` varchar(6) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `guides`
--

CREATE TABLE `guides` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `city` varchar(100) NOT NULL,
  `hourly_fee` decimal(10,2) NOT NULL,
  `languages` varchar(255) NOT NULL,
  `years_of_experience` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `guides`
--

INSERT INTO `guides` (`id`, `user_id`, `city`, `hourly_fee`, `languages`, `years_of_experience`, `created_at`, `updated_at`) VALUES
(13, 143, 'Sylhet', 200.00, 'Hindi,Bangla', 0, '2025-11-20 08:19:18', '2025-11-20 08:19:18'),
(14, 144, 'Sylhet', 300.00, 'bangla', 3, '2025-11-20 09:26:24', '2025-11-20 09:26:24');

-- --------------------------------------------------------

--
-- Table structure for table `guide_bookings`
--

CREATE TABLE `guide_bookings` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `guide_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `booking_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `duration_hours` int(11) NOT NULL,
  `total_fee` decimal(10,2) NOT NULL,
  `status` varchar(20) DEFAULT 'pending',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `phone` char(11) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password` varchar(255) NOT NULL,
  `photo` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `is_verified` tinyint(4) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `first_name`, `last_name`, `phone`, `email`, `password`, `photo`, `created_at`, `is_verified`) VALUES
(143, 'Ashraful', 'Islam', '01718111533', 'ashrafulislamkushimaro@gmail.com', '$2a$14$FgPieC3pn4yVXFLhHRMbiOjnj1HFSCMjAPaBLsJDWeYa.nXXefWZa', '', '2025-11-20 08:18:21', 1),
(144, 'Ashraful', 'Islam', '01718111533', 'ashrafulislamdarkeye@gmail.com', '$2a$14$0bim3OoD8pBSlLtopsjFFOu37qW6wETJpV7vxf8NiZjgZOi8QD6cO', '', '2025-11-20 09:24:14', 1),
(145, 'AIR', 'MAX', '01718111533', 'ashrafulialamraju@gmail.com', '$2a$14$EghuwYW1TOOXMXjyCyn2w.2G26.nVQITv6mLDczCxlrYCumtvx.nu', '', '2026-02-24 18:19:35', 1);

-- --------------------------------------------------------

--
-- Table structure for table `user_activity`
--

CREATE TABLE `user_activity` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `activity_type` varchar(100) NOT NULL,
  `description` text DEFAULT NULL,
  `page` varchar(150) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user_activity`
--

INSERT INTO `user_activity` (`id`, `user_id`, `activity_type`, `description`, `page`, `created_at`) VALUES
(3, 143, 'guides button', 'visiting guides page', '', '2025-11-20 14:19:06'),
(4, 143, 'settings button', 'visiting settings page', '', '2025-11-20 14:19:07'),
(5, 143, 'Guide Registration', 'User registered as a guide', '', '2025-11-20 14:19:18'),
(6, 143, 'guides button', 'visiting guides page', '', '2025-11-20 14:19:20'),
(7, 143, 'settings button', 'visiting settings page', '', '2025-11-20 14:58:02'),
(8, 143, 'planning button', 'visiting planning page', '', '2025-11-20 14:58:08'),
(9, 143, 'settings button', 'visiting settings page', '', '2025-11-20 14:58:09'),
(10, 143, 'settings button', 'visiting settings page', '', '2025-11-20 14:58:53'),
(11, 143, 'planning button', 'visiting planning page', '', '2025-11-20 14:59:04'),
(12, 143, 'settings button', 'visiting settings page', '', '2025-11-20 14:59:05'),
(13, 143, 'guides button', 'visiting guides page', '', '2025-11-20 15:06:23'),
(14, 143, 'settings button', 'visiting settings page', '', '2025-11-20 15:06:25'),
(16, 144, 'planning button', 'visiting planning page', '', '2025-11-20 15:24:49'),
(17, 144, 'guides button', 'visiting guides page', '', '2025-11-20 15:26:08'),
(18, 144, 'settings button', 'visiting settings page', '', '2025-11-20 15:26:10'),
(19, 144, 'Guide Registration', 'User registered as a guide', '', '2025-11-20 15:26:24'),
(20, 144, 'guides button', 'visiting guides page', '', '2025-11-20 15:26:26'),
(21, 144, 'monitor button', 'visiting monitor page', '', '2025-11-20 15:26:49'),
(22, 144, 'settings button', 'visiting settings page', '', '2025-11-20 15:27:27'),
(23, 145, 'settings button', 'visiting settings page', '', '2026-02-25 00:20:15'),
(24, 145, 'guides button', 'visiting guides page', '', '2026-02-25 00:20:21'),
(25, 145, 'settings button', 'visiting settings page', '', '2026-02-25 00:20:23'),
(26, 145, 'monitor button', 'visiting monitor page', '', '2026-02-25 00:20:24'),
(27, 145, 'settings button', 'visiting settings page', '', '2026-02-25 00:20:26');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `email_verification`
--
ALTER TABLE `email_verification`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `guides`
--
ALTER TABLE `guides`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `guide_bookings`
--
ALTER TABLE `guide_bookings`
  ADD PRIMARY KEY (`id`),
  ADD KEY `guide_id` (`guide_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `user_activity`
--
ALTER TABLE `user_activity`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `email_verification`
--
ALTER TABLE `email_verification`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=58;

--
-- AUTO_INCREMENT for table `guides`
--
ALTER TABLE `guides`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `guide_bookings`
--
ALTER TABLE `guide_bookings`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=146;

--
-- AUTO_INCREMENT for table `user_activity`
--
ALTER TABLE `user_activity`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `guides`
--
ALTER TABLE `guides`
  ADD CONSTRAINT `guides_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `guide_bookings`
--
ALTER TABLE `guide_bookings`
  ADD CONSTRAINT `guide_bookings_ibfk_1` FOREIGN KEY (`guide_id`) REFERENCES `guides` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `guide_bookings_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `user_activity`
--
ALTER TABLE `user_activity`
  ADD CONSTRAINT `user_activity_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
